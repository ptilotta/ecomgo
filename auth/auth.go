package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type tokenJSON struct {
	sub       string
	event_id  string
	token_use string
	scope     string
	auth_time int
	iss       string
	exp       int
	iat       int
	jti       string
	client_id string
	username  string
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
	userInfo, err := base64.StdEncoding.Strict().DecodeString(parts[1])
	if err != nil {
		fmt.Println("No se puede decodificar la parte del token:", err)
		return false, err, "No se puede decodificar la parte del token:"
	}
	// Aquí puedes hacer algo con la información de usuario, por ejemplo,
	// deserializarla en una estructura de datos
	fmt.Println("Información de usuario:", string(userInfo))

	// Serializamos el objeto userInfo dentro de la estructura correcta
	var jwt tokenJSON
	err = json.Unmarshal(userInfo, &jwt)
	if err != nil {
		fmt.Println("No se puede decodificar la estructura json:", err)
		return false, err, "No se puede decodificar la estructura json:"
	}

	fmt.Println("LA FECHA DEL TOKEN ES " + strconv.Itoa(jwt.exp))
	ahora := time.Now()
	tm := time.Unix(int64(jwt.exp), 0)

	if tm.Before(ahora) {
		fmt.Println(ahora.String() + " > " + tm.String())
		fmt.Println("Token expirado !")
		return false, err, "Token expirado !"
	}
	return true, nil, string(userInfo)
}
