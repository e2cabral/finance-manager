package helpers

import (
	"finance-manager/src/infra/environment"
	"finance-manager/src/infra/helpers/http"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func IsAuthenticated(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		env := environment.Env{}
		if r.Header["Authorization"] == nil {
			helpers.Forbidden(w, "Not authorized.")
			return
		}

		secret, err := env.GetVariable("SECRET_KEY")
		if err != nil {
			helpers.InternalServerError(w, err.Error())
			return
		}

		signString := []byte(fmt.Sprintf("%v", secret))

		token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("could not format")
			}
			return signString, nil
		})

		if err != nil {
			helpers.InternalServerError(w, err.Error())
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			handler.ServeHTTP(w, r)
			return
		}
		helpers.Forbidden(w, "Not authorized.")
	}
}
