package tools

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// FechaMySQL devuelve la fecha y hora actual en formato admitido por MySQL
func FechaMySQL() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

// Encrypt encripta el texto recibido y lo devuelve encriptado
func Encrypt(t string) (string, error) {
	text := []byte(t)

	hash, err := bcrypt.GenerateFromPassword(text, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
