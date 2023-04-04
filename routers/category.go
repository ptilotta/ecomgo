package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ptilotta/ecomgo/bd"
	"github.com/ptilotta/ecomgo/models"
)

/*InsertProduct es la funcion para crear en la BD el registro de producto */
func InsertCategory(body string, User string) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.CategName) == 0 {
		return 400, "Debe especificar el Nombre (Title) de la Categoria"
	}

	// Chequeamos que sea Admin quien hace la petición
	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertCategory(t)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar el registro de la categoría " + t.CategName + " > " + err.Error()
	}

	return 200, "{ CategID: " + strconv.Itoa(int(result)) + "}"
}

/*UpdateCategory es la funcion para modificar en la BD el registro de categoría */
func UpdateCategory(body string, User string, id int) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.CategName) == 0 && len(t.CategPath) == 0 {
		return 400, "Debe especificar Categ_Name o Categ_Path para actualizar"
	}

	// Chequeamos que sea Admin quien hace la petición
	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	t.CategID = id
	err2 := bd.UpdateCategory(t)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar el UPDATE de la categoria " + strconv.Itoa(t.CategID) + " > " + err.Error()
	}

	return 200, "Update OK"
}

/*DeleteCategory es la funcion para borrar en la BD el registro de categoría */
func DeleteCategory(body string, User string, id int) (int, string) {

	if id == 0 {
		return 400, "Debe especificar ID de la Categoría a borrar"
	}

	// Chequeamos que sea Admin quien hace la petición
	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	err := bd.DeleteCategory(id)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el DELETE de la categoria " + strconv.Itoa(id) + " > " + err.Error()
	}

	return 200, "Delete OK"
}

/*SelectCategory es la funcion para leer el registro de categoría */
func SelectCategories(body string, request events.APIGatewayV2HTTPRequest) (int, string) {
	var err error
	var CategId int
	var Slug string

	// Proceso los parámetros recibidos
	if len(request.QueryStringParameters["categId"]) > 0 {
		CategId, err = strconv.Atoi(request.QueryStringParameters["categId"])
		if err != nil {
			return 500, "Ocurrió un error al intentar convertir en entero al valor " + request.QueryStringParameters["categId"]
		}
	} else {
		if len(request.QueryStringParameters["slug"]) > 0 {
			Slug = request.QueryStringParameters["slug"]
		}
	}

	lista, err2 := bd.SelectCategories(CategId, Slug)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar capturar categoría/s > " + err2.Error()
	}

	Categ, err3 := json.Marshal(lista)

	if err3 != nil {
		return 400, "Ocurrió un error al intentar convertir en JSON los registros de Categorías"
	}

	return 200, string(Categ)
}
