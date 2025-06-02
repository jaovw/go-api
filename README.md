# 🧩 Go API
Uma API RESTful simples em **Golang** construída no intuido de aprendizado com foco em clareza, organização e boas praticas!  
Neste projeto, utilizo as bibliotecas **[gin-gonic](https://github.com/gin-gonic/gin)** para gerenciamento de rotas e **[gorm](https://gorm.io/)** como ORM, além de um banco de dados **SQLite** para persistência de dados leve e eficaz.

## 📦 Tecnologias Utilizadas

- **Go (Golang)** — Linguagem principal do projeto 🦦  
- **Gin Gonic** — Framework web leve e rápido ⚡  
- **GORM** — ORM intuitivo e robusto para Go 🛠️  
- **SQLite** — Banco de dados leve e embutido 💾  

## 🚀 Como Rodar o Projeto

### Pré-requisitos

- Go instalado (versão recomendada: **1.20+**)

### Passo a passo

```bash
# Clone o repositório
git clone https://github.com/jaovw/go-api.git
cd go-api

# Execute o projeto
go run main.go
```

> A API estará disponível em `http://localhost:8080`

## 📂 Estrutura do Projeto

```bash
go-api/
├── main.go
├── config/
│   └── config.go
├── handler/
│   └── handler.go
├── router/
│   └── router.go
│   └── routes.go
├── schema/
│   └── opening.go
└── go.mod
└── go.sum
```

## 🧭 Endpoints

Em construção 🚧  
A API já possui as bases para um CRUD completo. Em breve, endpoints documentados! 📘

## 📄 Documentation

Esta API foi desenvolvida com foco em clareza, desempenho e padronização. Para facilitar o entendimento e a integração com seus endpoints, utilizamos o **Swagger**, por meio da biblioteca **Swaggo**, uma das ferramentas mais populares para geração automática de documentação em APIs escritas em Go.

A documentação é gerada automaticamente a partir de comentários estruturados diretamente nos handlers da aplicação. Veja um exemplo:

```go
// @Summary Create opening
// @Description Create a new opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
```
> Para visualizar a documentação Swagger interativa, basta rodar o projeto e acessar a rota dedicada ao Swagger UI.


## 🪄 Makefile

Para facilitar a execução de comandos recorrentes durante o desenvolvimento, utilizamos um **Makefile**, permitindo interações rápidas via terminal com o comando `make`.

Abaixo estão os principais comandos disponíveis:

| Comando           | Descrição                                                         |
|-------------------|-------------------------------------------------------------------|
| `make` ou `make run` | Executa a aplicação diretamente usando `go run main.go`.         |
| `make build`      | Compila a aplicação e gera um executável com o nome `go_api`.     |
| `make test`       | Executa todos os testes presentes no projeto (`go test ./ ...`).  |
| `make docs`       | Gera a documentação Swagger utilizando o `swag init`.             |
| `make clean`      | Remove o executável gerado (`go_api.exe`) e o diretório `docs/`.  |

Esse processo torna o desenvolvimento mais produtivo e padronizado entre diferentes ambientes e usuários.

## 🤝 Contribuições

Sinta-se à vontade para abrir issues, pull requests ou dar sugestões! 💡  
Esse projeto é um playground para aprender, testar e crescer com Go!

## 📜 Licença

Este projeto está sob a licença [MIT](LICENSE).

Let’s build some APIs! 🛠️

