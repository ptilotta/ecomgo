package handlers

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ptilotta/ecomgo/auth"
	"github.com/ptilotta/ecomgo/routers"
)

/*Manejadores seteo mi puerto, el Handler y pongo a escuchar al Servidor */
func Manejadores(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {

	var User string
	fmt.Println("Voy a procesar " + path + " > " + method)
	param := request.PathParameters
	id := param["id"]
	idn, _ := strconv.Atoi(id)

	isOk, statusCode, message := validoAuthorization(path, method, id, idn, headers)
	if !isOk {
		return statusCode, message
	}

	User = message

	switch path[0:4] {
	case "user":
		return ProcesoUsers(body, path, method, User, id, request)
	case "prod":
		return ProcesoProduct(body, path, method, User, idn, request)
	case "stoc":
		return ProcesoStock(body, path, method, User, idn, request)
	case "addr":
		return ProcesoAddress(body, path, method, User, idn, request)
	case "cate":
		return ProcesoCategory(body, path, method, User, idn, request)
	case "orde":
		return ProcesoOrder(body, path, method, User, request)
	}

	return 400, "Method Invalid"
}

func validoAuthorization(path string, method string, id string, idn int, headers map[string]string) (bool, int, string) {
	if path == "user/me" || path == "users" || (path == "users/"+id && method == "DELETE") ||
		path == "order" || path == "orders" ||
		path == "address" ||
		(path == "product" && method != "GET") ||
		(path == "product/"+strconv.Itoa(idn) && (method == "PUT" || method == "DELETE")) ||
		path == "stock/"+strconv.Itoa(idn) ||
		(path == "categories" && method != "GET") || (path == "category" && method != "GET") ||
		(path == "category/"+strconv.Itoa(idn) && (method == "PUT" || method == "DELETE")) {

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

func ProcesoUsers(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {

	if path == "user/me" {
		switch method {
		case "PUT":
			return routers.UpdateUser(body, user)
		case "GET":
			return routers.SelectUser(body, user)
		}
	}
	if path == "users" {
		if method == "GET" {
			return routers.SelectUsers(body, user, request)
		}
	}
	if path == "users/"+id && method == "DELETE" {
		return routers.DeleteUser(user, id)
	}
	return 400, "Method Invalid"
}

func ProcesoProduct(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	switch method {
	case "POST":
		return routers.InsertProduct(body, user)
	case "GET":
		return routers.SelectProduct(body, request)
	case "PUT":
		return routers.UpdateProduct(body, user, id)
	case "DELETE":
		return routers.DeleteProduct(user, id)
	}
	return 400, "Method Invalid"
}

func ProcesoStock(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	if method == "PUT" {
		return routers.UpdateStock(body, user, id)
	}
	return 400, "Method Invalid"
}

func ProcesoAddress(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	switch method {
	case "POST":
		return routers.InsertAddress(body, user)
	case "GET":
		return routers.SelectAddresses(user)
	case "PUT":
		return routers.UpdateAddress(body, user, id)
	case "DELETE":
		return routers.DeleteAddress(user, id)
	}
	return 400, "Method Invalid"
}

func ProcesoCategory(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	switch method {
	case "POST":
		return routers.InsertCategory(body, user)
	case "PUT":
		return routers.UpdateCategory(body, user, id)
	case "DELETE":
		return routers.DeleteCategory(body, user, id)
	case "GET":
		return routers.SelectCategories(body, request)
	}
	return 400, "Method Invalid"
}

func ProcesoOrder(body string, path string, method string, user string, request events.APIGatewayV2HTTPRequest) (int, string) {

	if path == "orders" && method == "GET" {
		return routers.SelectOrders(user, request)
	}

	switch method {
	case "POST":
		return routers.InsertOrder(body, user)
	case "GET":
		return routers.SelectOrder(body, request)
	}
	return 400, "Method Invalid"
}
