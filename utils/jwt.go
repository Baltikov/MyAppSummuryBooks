package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"testapi/pkg/loger"
	"time"
)

const SecretKey = "your-256-bit-secret"

func GenerateJwt(email string, userID int64) (string, error) {
	// Создание нового токена с использованием указанных claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	// Подписание токена с секретным ключом
	stringToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		loger.Logrus.Error("Error signing token: %v", err)
		loger.Logrus.Trace(err.Error())
		return "", err
	}

	return stringToken, nil
}

// Checkjwt проверка каким методом кодирирования был создан ключ

func CheckJwt(tokenString string) (float64, error) {
	tokenString = strings.TrimSpace(tokenString)
	fmt.Println("tokenString is", tokenString)
	// Парсинг и валидация токена
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверка метода подписи- Метод которым подписали ключ - этот SigningMethodHMAC?
		// если нет, то выдает ошибку
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	// Фунция возращает ключ в виде среза байтов, но фукнция не может распрасить
	// не получаю на вход байты, для парса
	fmt.Println("token is", token)

	if err != nil {
		return 0, err
	}

	// Проверка валидности токена
	// Проверяе соотвествуют ли Claim внутри токена типу по которому они были записаны wt.MapClaims
	// и проверика на то, что токена подписан секретным ключом, который юыл ранее
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//email := claims["email"].(string)
		userID := claims["userId"].(float64)
		return userID, nil
	} else {
		return 0, fmt.Errorf("Invalid token")
	}
	return 0, nil
}
