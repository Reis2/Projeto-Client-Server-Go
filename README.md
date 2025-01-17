# Projeto: Cotação do Dólar

## Objetivos

O objetivo do projeto é obter a cotação atual do dólar em relação ao real brasileiro no site https://economia.awesomeapi.com.br/json/last/USD-BRL por meio da requisição HTTP em linguagem GO e salvar os mesmos dados em dólar em um arquivo de banco de dados (.db) chamado de "cotacoes" presente na pasta "server". E, por último, entregar o valor atual e atualizado a um arquivo criado chamado "cotacao" em "/client".

## Ferramentas

- [Go](https://golang.org/doc/install) instalado na versão 1.16 ou superior.

## Funcionalidade

- Realiza uma requisição no código **server.go** da cotacao atual do dólar (USD) em realção ao real (BRL) no site [economia.awesomeapi](https://economia.awesomeapi.com.br/json/last/USD-BRL).
- No código principal, **client.go**, o resultado obtido pelo server é requisitado pelo método GET do localhost (porta 8080).
- Após a execução dos programas, é possível observar que os arquivos cotacao.txt e cotacoes.db são criados ou atualizados com a cotação atual do dólar.

## Como Executar

1. **Clone o repositório:**
   ```console
   git clone thttps://github.com/Reis2/Projeto-Client-Server-Go
   cd Projeto-Client-Server-Go

2. **Executando programa**
   
   2.1 Abrindo um terminal dedicado a uma abertura do servidor:
      ```console
      cd server
      go run server.go
      ```
   2.1 Abrindo um terminal para executar o programa client.go:
      ```console
      cd client
      go run client.go
      ```

**Notas:**

Este `README.md` fornece uma visão geral clara do propósito do projeto, instruções de execução e referências úteis para usuários e colaboradores.



      


   

   
