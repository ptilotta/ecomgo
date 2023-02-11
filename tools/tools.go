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
