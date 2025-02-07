package repositories

import (
	configs "GOUSERAPI/config"
	"database/sql"
	"gouserapi/models"
)

func CreateUser(user *models.User) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err := configs.DB.Exec(query, user.Name, user.Email)
	return err
}

func GetUserByID(id string) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, email FROM users WHERE id = ?"
	err := config.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		return nil, err
	}
	return &user, err
}

func CheckUserExists(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = ? LIMIT 1)"
	err := configs.DB.QueryRow(query, email).Scan(&exists)
	return exists, err
}
