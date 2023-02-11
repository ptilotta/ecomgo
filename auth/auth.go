package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	Sub       string
	Event_id  string
	Token_use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Jti       string
	Client_id string
	Username  string
}

func ValidoToken(token string) (bool, error, string) {
	// Separamos el token en tres partes, separadas por "."
	parts := strings.Split(token, ".")

	// Validamos que tengan 3 partes
	if len(parts) != 3 {
		fmt.Println("El token no es válido.")
		return false, nil, "El token no es válido."
	}

	// La segunda parte contiene la información de usuario codificada
	// como una estructura JSON. Debemos decodificarla.
	// En este ejemplo, usamos base64.StdEncoding.DecodeString para
	// decodificar la parte en una cadena.
	userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("No se puede decodificar la parte del token:", err)
		return false, err, "No se puede decodificar la parte del token:"
	}

	// Aquí puedes hacer algo con la información de usuario, por ejemplo,
	// deserializarla en una estructura de datos
	// Serializamos el objeto userInfo dentro de la estructura correcta
	var tkj TokenJSON
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		fmt.Println("No se puede decodificar la estructura json:", err)
		return false, err, "No se puede decodificar la estructura json:"
	}

	ahora := time.Now()
	tm := time.Unix(int64(tkj.Exp), 0)

	if tm.Before(ahora) {
		fmt.Println("Fecha expiración token = " + tm.String())
		fmt.Println("Token expirado !")
		return false, err, "Token expirado !"
	}
	return true, nil, string(tkj.Username)
}
