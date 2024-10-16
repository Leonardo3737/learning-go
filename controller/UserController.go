package controller

import (
	"api/models"
	"api/service"
	"api/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func UserController(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	sid := strings.TrimPrefix(r.URL.Path, "/user/")
	id, _ := strconv.Atoi(sid)

	fmt.Println(id)

	switch {
	case r.Method == "GET" && id > 0:
		handleGetUserById(w, db, id)
	case r.Method == "GET":
		handleGetUsers(w, db)
	case r.Method == "POST":
		handlePostUser(w, r, db)
	case r.Method == "PUT" && id > 0:
		handlePutUser(w, r, db, id)
	case r.Method == "DELETE" && id > 0:
		handleDeleteUser(w, db, id)
	default:
		utils.MakeErrorResponse(w, http.StatusMethodNotAllowed, "Método não permitido")
	}
}

func handleGetUsers(w http.ResponseWriter, db *sql.DB) {
	users, getErr := service.GetUser(db)

	if getErr != nil {
		utils.MakeErrorResponse(w, http.StatusInternalServerError, "Erro ao buscar usuários.")
		return
	}
	utils.MakeJsonResponse(w, http.StatusOK, users)
}

func handleGetUserById(w http.ResponseWriter, db *sql.DB, id int) {
	user, getErr := service.GetUserById(db, id)

	if getErr != nil {
		if getErr == sql.ErrNoRows {
			utils.MakeErrorResponse(w, http.StatusNotFound, "usuário não encontrado.")
		} else {
			utils.MakeErrorResponse(w, http.StatusInternalServerError, "Erro ao buscar usuário.")
		}
		return
	}
	utils.MakeJsonResponse(w, http.StatusOK, user)
}

func handlePostUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var newUser models.User

	jsonErr := json.NewDecoder(r.Body).Decode(&newUser)
	if jsonErr != nil {
		utils.MakeErrorResponse(w, http.StatusBadRequest, "Erro ao gerar usuário.")
		return
	}
	defer r.Body.Close()

	validationErr := validate.Struct(newUser)
	if validationErr != nil {
		utils.MakeErrorResponse(w, http.StatusBadRequest, "Dados do usuário são inválidos.")
		return
	}

	createErr := service.CreateUser(db, &newUser)
	if createErr != nil {
		utils.MakeErrorResponse(w, http.StatusInternalServerError, "Erro ao criar usuário.")
		return
	}
	utils.MakeJsonResponse(w, http.StatusCreated, newUser)
}

func handlePutUser(w http.ResponseWriter, r *http.Request, db *sql.DB, id int) {
	_, getErr := service.GetUserById(db, id)

	if getErr != nil {
		if getErr == sql.ErrNoRows {
			utils.MakeErrorResponse(w, http.StatusNotFound, "usuário não encontrado.")
		} else {
			utils.MakeErrorResponse(w, http.StatusInternalServerError, "Erro ao buscar usuário.")
		}
		return
	}

	var user models.User

	jsonErr := json.NewDecoder(r.Body).Decode(&user)

	if jsonErr != nil {
		utils.MakeErrorResponse(w, http.StatusBadRequest, "Erro ao gerar usuário.")
		return
	}
	defer r.Body.Close()

	validationErr := validate.Struct(user)
	if validationErr != nil {
		utils.MakeErrorResponse(w, http.StatusBadRequest, "Dados do usuário são inválidos.")
		return
	}

	updateErr := service.UpdateUser(db, user, id)
	if updateErr != nil {
		utils.MakeErrorResponse(w, http.StatusInternalServerError, "Erro ao autualizar usuário.")
		return
	}
	utils.MakeJsonResponse(w, http.StatusNoContent, nil)
}

func handleDeleteUser(w http.ResponseWriter, db *sql.DB, id int) {

	_, getErr := service.GetUserById(db, id)

	if getErr != nil {
		if getErr == sql.ErrNoRows {
			utils.MakeErrorResponse(w, http.StatusNotFound, "usuário não encontrado.")
		} else {
			utils.MakeErrorResponse(w, http.StatusInternalServerError, "Erro ao buscar usuário.")
		}
		return
	}

	deleteErr := service.DeleteUser(db, id)
	if deleteErr != nil {
		utils.MakeErrorResponse(w, http.StatusInternalServerError, "Erro ao deletar usuário.")
		return
	}
	utils.MakeJsonResponse(w, http.StatusNoContent, nil)
}
