package tools

import (
	"fmt"
	"time"
)

var Email string
var Expirate int64

// FechaMySQL devuelve la fecha y hora actual en formato admitido por MySQL
func FechaMySQL() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}
