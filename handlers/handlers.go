package handlers

import (
	"fmt"

	"github.com/ptilotta/ecomgo/routers"
	"github.com/ptilotta/ecomgo/tools"
)

/*Manejadores seteo mi puerto, el Handler y pongo a escuchar al Servidor */
func Manejadores(path string, method string, body string, headers map[string]string, userPoolId string, region string) (int, string) {

	fmt.Println("event.Path = " + path + " - event.HTTPMethod = " + method)

	switch path {
	case "/user/me":
		for key, value := range headers {
			fmt.Println("Key: ", key, "Value: ", value)
		}
		err := tools.ValidateJWT(headers["authorization"], userPoolId, region)
		if err != nil {
			fmt.Println("Error : JWT Inválido")
			return 400, "Error : JWT Inválido"
		}

		if method == "POST" {
			fmt.Println("Voy al routers.UpdateUser(body)")
			return routers.UpdateUser(body)
		}
	}

	return 200, "Todo OK"
}
