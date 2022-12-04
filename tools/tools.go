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

/* URL para probar el token de Cognito
https://cognito-idp.{region}.amazonaws.com/{userPoolId}/.well-known/jwks.json
User Pool ID = us-east-1_ln0bNuRBv
Region = us-east-1
https://cognito-idp.us-east-1.amazonaws.com/us-east-1_ln0bNuRBv/.well-known/jwks.json
*/
