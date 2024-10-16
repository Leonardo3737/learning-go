package main

import (
	mydb "api/mydb"
	"api/routes"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const DBConnetion string = "root:030131@tcp(localhost:3306)/"

func main() {
	db, err := sql.Open("mysql", DBConnetion)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	mydb.DBInit(db)

	fmt.Println("Conectado ao banco de dados MySQL com sucesso!")

	routes.RoutesInit(db)

	fmt.Println("Rodando na porta 3000")

	http.ListenAndServe(":3000", nil)
}
