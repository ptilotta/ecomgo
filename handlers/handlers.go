package handlers

import (
	"fmt"

	"github.com/ptilotta/ecomgo/routers"
)

/*Manejadores seteo mi puerto, el Handler y pongo a escuchar al Servidor */
func Manejadores(path string, method string, body string, headers map[string]string) (int, string) {

	fmt.Println("event.Path = " + path + " - event.HTTPMethod = " + method)

	switch path {
	case "/default/ecommerce/user/me":
		switch method {
		case "POST":
			fmt.Println("Voy al routers.UpdateUser(body)")
			return routers.UpdateUser(body)
		case "GET":
			fmt.Println("Voy al routers.SelectUser(body)")
			return routers.SelectUser(body)
		case "DELETE":
			fmt.Println("Voy al routers.SelectUser(body)")
			return routers.SelectUser(body)
		}
	case "/default/ecommerce/users":
		if method == "GET" {
			fmt.Println("Voy al routers.SelectUsers(body)")
			return routers.SelectUsers(body)
		}
	}

	return 200, "Todo OK"
}
