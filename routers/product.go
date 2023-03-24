package routers

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ptilotta/ecomgo/bd"
	"github.com/ptilotta/ecomgo/models"
)

/*InsertProduct es la funcion para crear en la BD el registro de producto */
func InsertProduct(body string, User string) (int, string) {
	var t models.Product
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.ProdTitle) == 0 {
		return 400, "Debe especificar el Nombre (Title) del Producto"
	}

	// Chequeamos que sea Admin quien hace la petición
	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertProduct(t)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar el registro del producto " + t.ProdTitle + " > " + err.Error()
	}

	return 200, "{ ProductID: " + strconv.Itoa(int(result)) + "}"
}

/*SelectProduct es la funcion para crear en la BD el registro de producto */
func SelectProduct(body string, request events.APIGatewayV2HTTPRequest) (int, string) {
	var t models.Product
	var err error
	var page, pageSize int
	var orderType, orderField string

	param := request.QueryStringParameters

	page, _ = strconv.Atoi(param["page"])
	pageSize, _ = strconv.Atoi(param["pageSize"])
	orderType = param["orderType"]   // D = Desc, A or null = Asc
	orderField = param["orderField"] // 'I' or null, 'T' Title,   'D' Description, 'F' Created At,
	// 'P' Price,   'C' CategId, 'S' Stock

	if !strings.Contains("ITDFPCS", orderType) {
		orderField = ""
	}

	var choice string
	if len(param["prodId"]) > 0 {
		choice = "P"
		t.ProdID, _ = strconv.Atoi(param["prodId"])
	}
	if len(param["search"]) > 0 {
		choice = "S"
		t.ProdSearch = param["search"]
	}
	if len(param["categId"]) > 0 {
		choice = "C"
		t.ProdCategId, _ = strconv.Atoi(param["categId"])
	}

	if len(choice) == 0 {
		return 400, "Debe especificar el ID del Producto o el parámetro 'search' o el Id de una categoría"
	}

	result, err2 := bd.SelectProduct(t, choice, page, pageSize, orderType, orderField)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar capturar los resultados de la búsqueda de tipo '" + choice + "' en productos > " + err.Error()
	}

	Product, err3 := json.Marshal(result)
	if err3 != nil {
		return 400, "Ocurrió un error al intentar convertir en JSON el registro de Producto"
	}

	return 200, string(Product)
}

/*UpdateProduct es la funcion para modificar en la BD el registro de producto */
func UpdateProduct(body string, User string) (int, string) {
	var t models.Product
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	// Chequeamos que sea Admin quien hace la petición
	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	err2 := bd.UpdateProduct(t)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar el UPDATE del producto " + strconv.Itoa(t.ProdID) + " > " + err.Error()
	}

	return 200, "Update OK"
}

/*DeleteProduct es la funcion para borrar en la BD el registro de producto */
func DeleteProduct(body string, User string) (int, string) {
	var t models.Product
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	// Chequeamos que sea Admin quien hace la petición
	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	err2 := bd.DeleteProduct(t)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar el Delete del producto " + strconv.Itoa(t.ProdID) + " > " + err.Error()
	}

	return 200, "Delete OK"
}

/*UpdateStock es la funcion para actualizar el Stock de un producto */
func UpdateStock(body string, User string) (int, string) {
	var t models.Product
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	// Chequeamos que sea Admin quien hace la petición
	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	err2 := bd.UpdateStock(t)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar el Update del Stock del producto " + strconv.Itoa(t.ProdID) + " > " + err.Error()
	}

	return 200, "Update Stock OK"
}
