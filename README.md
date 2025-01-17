# Projeto: Cotação do Dólar
O objetivo do projeto é obter a cotação atual do dólar em relação ao real brasileiro no site https://economia.awesomeapi.com.br/json/last/USD-BRL por meio da requisição HTTP em linguagem GO e salvar os mesmos dados em dólar em um arquivo de banco de dados (.db) chamado de "cotacoes.db" presente na pasta "server". E, por último, entregar o valor atual e atualizado a um arquivo criado chamado "cotacao.txt" em "/client".

##Ferramentas
- [Go](https://golang.org/doc/install) instalado na versão 1.16 ou superior.

##Funcionalidade
- Realiza uma requisição no código **server.go** da cotacao atual do dólar (USD) em realção ao real (BRL) no site [economia.awesomeapi](https://economia.awesomeapi.com.br/json/last/USD-BRL).
- No código **client.go** o resultado obtido pelo server é requisitado pelo método GET do localhost (porta 8080).
- Ao averiguar possíivel observar a presença dos respectivos dados no banco de dados e no arquivo de texto.

## Como Executar

1. **Clone o repositório:**
   ```terminal
   git clone https://github.com/seu-usuario/seu-repositorio.git
   cd seu-repositorio
