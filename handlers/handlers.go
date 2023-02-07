package handlers

import (
	"fmt"

	"github.com/ptilotta/ecomgo/auth"
	"github.com/ptilotta/ecomgo/routers"
)

/*Manejadores seteo mi puerto, el Handler y pongo a escuchar al Servidor */
func Manejadores(path string, method string, body string, headers map[string]string) (int, string) {

	fmt.Println("event.Path = " + path + " - event.HTTPMethod = " + method)

	var User string

	if path == "/default/ecommerce/user/me" ||
		path == "/default/ecommerce/users" ||
		(path == "/default/ecommerce/product" && method == "POST") {

		token := headers["Authorization"]

		if len(token) == 0 {
			return 401, "Token requerido"
		}

		todoOK, err2, msg := auth.ValidoToken(token)

		if !todoOK {
			if err2 != nil {
				fmt.Println("Error en el token " + err2.Error())
				return 401, err2.Error()
			} else {
				fmt.Println("Error en el token " + msg)
				return 401, msg
			}
		} else {
			fmt.Println("Token OK")
			User = msg
		}
	}

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
			fmt.Println("Voy al routers.DeleteUser(body)")
			return routers.DeleteUser(body)
		}
	case "/default/ecommerce/users":
		if method == "GET" {
			fmt.Println("Voy al routers.SelectUsers(body)")
			return routers.SelectUsers(body)
		}

	case "/default/ecommerce/product":
		if method == "POST" {
			fmt.Println("Voy al routers.SelectUsers(body)")
			return routers.InsertProduct(body)
		}
	}

	return 200, "Todo OK"
}
