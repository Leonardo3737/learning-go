package routes

import (
	controller "api/controller"
	"database/sql"
	"net/http"
)

func RoutesInit(db *sql.DB) {
	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		controller.UserController(w, r, db)
	})
}
