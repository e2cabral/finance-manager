package auth

import (
	"finance-manager/src/infra/environment"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateJWTToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	env := environment.Env{}
	var err error

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	if err != nil {
		return "", err
	}

	secret, err := env.GetVariable("SECRET_KEY")
	if err != nil {
		return "", err
	}

	convertedString := []byte(fmt.Sprintf("%v", secret))

	tokenString, err := token.SignedString(convertedString)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
