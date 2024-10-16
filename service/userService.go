package service

import (
	"api/models"
	"database/sql"
)

func CreateUser(db *sql.DB, user *models.User) error {
	stmt, _ := db.Prepare(`
	INSERT INTO user (
	name
	) values (
	 ?
	 )
	 `)

	res, err := stmt.Exec(user.Name)

	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	user.ID = int(id)
	return nil
}

func GetUser(db *sql.DB) ([]models.User, error) {
	row, err := db.Query("SELECT * FROM user")

	if err != nil {
		return nil, err
	}

	var users []models.User

	for row.Next() {
		var user models.User
		row.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}

	return users, nil
}

func GetUserById(db *sql.DB, id int) (models.User, error) {
	var user models.User

	err := db.QueryRow("SELECT id_user, name FROM user WHERE id_user = ?", id).Scan(&user.ID, &user.Name)

	if err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUser(db *sql.DB, user models.User, id int) error {
	stmt, _ := db.Prepare(`
	UPDATE user SET name = ? WHERE id_user = ?
	`)
	_, err := stmt.Exec(user.Name, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(db *sql.DB, id int) error {
	stmt, _ := db.Prepare(`
	DELETE FROM user WHERE id_user = ?
	`)
	_, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
