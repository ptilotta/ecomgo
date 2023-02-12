package handlers

import (
	"fmt"

	"github.com/ptilotta/ecomgo/auth"
	"github.com/ptilotta/ecomgo/routers"
)

/*Manejadores seteo mi puerto, el Handler y pongo a escuchar al Servidor */
func Manejadores(path string, method string, body string, headers map[string]string) (int, string) {

	var User string

	isOk, statusCode, message := validoAuthorization(path, method, headers)
	if !isOk {
		return statusCode, message
	}

	User = message

	switch path {
	case "/default/ecommerce/user/me":
		return UserCRUD(body, path, method, User)
	case "/default/ecommerce/users":
		if method == "GET" {
			fmt.Println("Voy al routers.SelectUsers(body, User)")
			return routers.SelectUsers(body, User)
		}
	case "/default/ecommerce/product":
		return ProductCRUD(body, path, method, User)
	case "/default/ecommerce/stock":
		if method == "PUT" {
			return routers.UpdateStock(body, User)
		}
	case "/default/ecommerce/category":
		return CategoryCRUD(body, path, method, User)
	}

	return 200, "Todo OK"
}

func validoAuthorization(path string, method string, headers map[string]string) (bool, int, string) {
	if path == "/default/ecommerce/user/me" ||
		path == "/default/ecommerce/users" ||
		path == "/default/ecommerce/stock" ||
		path == "/default/ecommerce/category" ||
		(path == "/default/ecommerce/product" && method != "GET") {

		fmt.Println(headers)
		token := headers["authorization"]

		if len(token) == 0 {
			return false, 401, "Token requerido"
		}

		todoOK, err2, msg := auth.ValidoToken(token)

		if !todoOK {
			if err2 != nil {
				fmt.Println("Error en el token " + err2.Error())
				return false, 401, err2.Error()
			} else {
				fmt.Println("Error en el token " + msg)
				return false, 401, msg
			}
		} else {
			fmt.Println("Token OK")
			return true, 200, msg
		}
	}

	return true, 200, ""
}

func UserCRUD(body string, path string, method string, user string) (int, string) {
	fmt.Println("Voy a procesar " + path + " > " + method + " para el user " + user)
	switch method {
	case "POST":
		return routers.UpdateUser(body, user)
	case "GET":
		return routers.SelectUser(body, user)
	case "DELETE":
		return routers.DeleteUser(body, user)
	}
	return 400, "Method Invalid"
}

func ProductCRUD(body string, path string, method string, user string) (int, string) {
	fmt.Println("Voy a procesar " + path + " > " + method)
	switch method {
	case "POST":
		return routers.InsertProduct(body, user)
	case "GET":
		return routers.SelectProduct(body)
	case "PUT":
		return routers.UpdateProduct(body, user)
	case "DELETE":
		return routers.DeleteProduct(body, user)
	}
	return 400, "Method Invalid"
}

func CategoryCRUD(body string, path string, method string, user string) (int, string) {
	fmt.Println("Voy a procesar " + path + " > " + method)
	switch method {
	case "POST":
		return routers.InsertCategory(body, user)
	case "PUT":
		return routers.UpdateCategory(body, user)
	case "GET":
		return routers.SelectCategory(body)
	}
	return 400, "Method Invalid"
}
