package tools

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var Email string
var Expirate int64

type MyCustomClaims struct {
	User_name string `json:"user_name"`
	jwt.StandardClaims
}

// FechaMySQL devuelve la fecha y hora actual en formato admitido por MySQL
func FechaMySQL() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

// EscapeString quita comillas simples y dobles de un String
func EscapeString(t string) string {
	desc := strings.ReplaceAll(t, "'", "")
	desc = strings.ReplaceAll(desc, "\"", "")
	return desc
}

func ValidoJWT(t string) (bool, error, string) {
	var mySecret = "8Xi/PEzDz4P6m9cRMLGZ7ilcxBHIdZfnEgEpw/q4IwA="

	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return mySecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user_name := fmt.Sprintf("%s", claims["User_name"])
		return true, nil, user_name
	} else {
		return false, err, ""
	}
}
