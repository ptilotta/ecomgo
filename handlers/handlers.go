package handlers

import (
	"fmt"

	"github.com/ptilotta/ecomgo/routers"
)

/*Manejadores seteo mi puerto, el Handler y pongo a escuchar al Servidor */
func Manejadores(path string, method string, body string, headers map[string]string, userPoolId string, region string) (int, string) {

	fmt.Println("event.Path = " + path + " - event.HTTPMethod = " + method)

	switch path {
	case "/user/me":
		if method == "POST" {
			fmt.Println("Voy al routers.UpdateUser(body)")
			return routers.UpdateUser(body)
		}
	}

	return 200, "Todo OK"
}
