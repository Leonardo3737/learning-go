# API REST em Go - Projeto de Estudo

  

Este é um projeto simples de uma API REST desenvolvida em Go, criado com fins de estudo. A API implementa operações básicas de CRUD (Create, Read, Update, Delete) para uma única entidade, chamada `User`.

  

## Pré-requisitos

  

- Go 1.18 ou superior instalado em sua máquina.

- [Docker](https://www.docker.com/) (opcional, para rodar o banco de dados MySQL).

  

## Configuração do Ambiente

1. Clone o repositório:
```
git clone https://github.com/Leonardo3737/learning-go

cd learning-go
```
2. Baixe as dependências do projeto:
```
go mod tidy
```
3. Configure a conexão com o banco de dados (mysql):
 - Altere a string de conexão (constante DBConnetion no aquivo main.go, linha 14) para sua conexão

4. Rode o projeto:
```
go run main.go
```

## Rotas
  - GET http://localhost:3000/user
```
Response:
[
  {
    "id": 1,
    "name": "your name"
  }
]
```

- POST http://localhost:3000/user
- PUT http://localhost:3000/user/{id}
```
Body:
{
  "name": "your name"  
}
```
- DELETE http://localhost:3000/user/{id}