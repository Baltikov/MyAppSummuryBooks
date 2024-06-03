package model

import (
	"errors"
	"testapi/bd"
	"testapi/pkg/loger"
	"testapi/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Save(user *User) error {
	queryCreate := `INSERT INTO users (email, password) VALUES (?, ?)`
	hashPass, err := utils.Hash(user.Password)
	if err != nil {
		loger.Logrus.Error(err)
	}
	result, err := bd.DB.Exec(queryCreate, &user.Email, &hashPass)
	if err != nil {
		loger.Logrus.Error(err)
		return err
	}
	user.ID, err = result.LastInsertId()
	return nil
}

func CheckUser(user *User) error {
	query := `SELECT id, password users WHERE email = ? LIMIT 1`
	row, err := bd.DB.Query(query, user.Email)
	if err != nil {
		loger.Logrus.Error(err)
		return err
	}
	var checkPass string
	// мы проверяем существует ли пользователь
	// Если да захешим пароль и зименим айди
	err = row.Scan(&user.ID, &checkPass)
	if err != nil {
		loger.Logrus.Error(err)
		return err
	}
	isValidPassword := utils.CheckPasswordHash(user.Password, checkPass)

	if !isValidPassword {
		return errors.New("invalid password")
	}
	defer row.Close()

	return nil
}
