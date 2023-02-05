package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type tokenJSON struct {
	sub       string `json:"sub"`
	event_id  string `json:"event_id"`
	token_use string `json:"token_use"`
	scope     string `json:"scope"`
	auth_time int    `json:"auth_time"`
	iss       string `json:"iss"`
	exp       int    `json:"exp"`
	iat       int    `json:"iat"`
	jti       string `json:"jti"`
	client_id string `json:"client_id"`
	username  string `json:"username"`
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
	fmt.Println("Información de usuario:", string(userInfo))

	// Serializamos el objeto userInfo dentro de la estructura correcta
	var tkj tokenJSON
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		fmt.Println("No se puede decodificar la estructura json:", err)
		return false, err, "No se puede decodificar la estructura json:"
	}

	fmt.Println("tkj.sub = " + tkj.sub)
	ahora := time.Now()
	tm := time.Unix(int64(tkj.exp), 0)

	if tm.Before(ahora) {
		fmt.Println(ahora.String() + " > " + tm.String())
		fmt.Println("Token expirado !")
		return false, err, "Token expirado !"
	}
	return true, nil, string(userInfo)
}
