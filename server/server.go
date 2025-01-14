package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Estrutura para armazenar a resposta da API
type APIResponse struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

// Inicializa o banco de dados SQLite
func initDB() (*sql.DB, error) {
	//println("Tudo certo")
	db, err := sql.Open("sqlite3", "./cotacoes.db")
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, fmt.Errorf("falha ao abrir o banco de dados: objeto db é nulo")
	}

	// Criação da tabela, se não existir
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS cotacoes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bid TEXT
	);`

	if _, err := db.Exec(createTableQuery); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatalf("Erro ao inicializar o banco de dados: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		// Contexto com timeout para consultar a API externa
		ctxAPI, cancelAPI := context.WithTimeout(context.Background(), 1000*time.Millisecond)
		defer cancelAPI()

		req, err := http.NewRequestWithContext(ctxAPI, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				fmt.Println("Tempo de requisição à API externa excedido")
			} else {
				fmt.Printf("Erro ao fazer a requisição: %v\n", err)
			}
			return
		}
		defer resp.Body.Close()

		// Verificar se o status da resposta é OK (200)
		if resp.StatusCode != http.StatusOK {
			fmt.Fprintf(os.Stderr, "Erro: Status da resposta não é OK: %v\n", resp.Status)
			return
		}

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler o corpo da resposta: %v\n", err)
			return
		}

		var data APIResponse
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao parse da resposta resposta: %v\n", err)
		}

		// Contexto com timeout para persistir no banco
		ctxDB, cancelDB := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancelDB()

		// Insere os dados no banco
		insertQuery := `
		INSERT INTO cotacoes (bid)
		VALUES (?);`
		_, err = db.ExecContext(ctxDB, insertQuery,
			data.USDBRL.Bid,
		)

		if err != nil {
			http.Error(w, "Erro ao salvar cotação no banco", http.StatusInternalServerError)
			return
		}

		// Retornando a cotação do dólar pelo "data" no formato JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data.USDBRL.Bid)

	})

	http.ListenAndServe(":8080", nil)
}
