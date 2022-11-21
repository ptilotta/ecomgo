package tools

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"fmt"
	"io"
	"time"
)

// FechaMySQL devuelve la fecha y hora actual en formato admitido por MySQL
func FechaMySQL() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

// Encrypt encripta el texto recibido y lo devuelve encriptado
func Encrypt(t string) ([]byte, error) {
	text := []byte(t)
	key := []byte("acomgo-hecho-en-go-y-reactnative")

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err2 := cipher.NewGCM(c)
	if err2 != nil {
		return nil, err2
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, text, nil), nil
}
