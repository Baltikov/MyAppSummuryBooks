package utils

import (
	"golang.org/x/crypto/bcrypt"
	"testapi/pkg/loger"
)

func Hash(s string) (string, error) {
	passord, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		loger.Logrus.Error(err)
		return "", nil
	}
	return string(passord), nil
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		loger.Logrus.Error(err)
		return false
	}
	return true
}
