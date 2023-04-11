package tools

import (
	"fmt"
	"strconv"
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

func ArmoSentencia(s string, fieldName string, typeField string, ValueN int, ValueF float64, ValueS string) string {
	if (typeField == "S" && len(ValueS) == 0) || (typeField == "F" && ValueF == 0) || (typeField == "N" && ValueN == 0) {
		return s
	}

	if !strings.HasSuffix(s, "SET ") {
		s += ", "
	}

	switch typeField {
	case "S":
		s += fieldName + " = '" + EscapeString(ValueS) + "'"
	case "N":
		s += fieldName + " = " + strconv.Itoa(ValueN)
	case "F":
		s += fieldName + " = " + strconv.FormatFloat(ValueF, 'e', -1, 64)
	}

	return s
}
