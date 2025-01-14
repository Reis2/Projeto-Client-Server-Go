package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// Definir um contexto com timeout de 300ms
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	// Criar uma requisição HTTP com o contexto
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Printf("Erro ao criar a requisição: %v\n", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("Tempo de requisição ao servidor excedido")
		} else {
			fmt.Printf("Erro ao fazer a requisição: %v\n", err)
		}
		return
	}
	defer resp.Body.Close()

	dolar, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v\n", err)
	}

	//fmt.Println(string(body))

	/* 	var dolar USDBRL
	   	err = json.Unmarshal(body, &dolar)
	   	if err != nil {
	   		fmt.Fprintf(os.Stderr, "Erro ao parse da resposta resposta: %v\n", err)
	   	}
	*/
	// Exibindo os dados da cotação
	fmt.Printf("Valor da cotação em dólar: %s\n", dolar)

	file, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar o arquivo: %v\n", err)
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("Dólar: %s", dolar))
}
