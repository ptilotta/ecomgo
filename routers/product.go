package routers

import (
	"encoding/json"
	"strconv"

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
func SelectProduct(body string) (int, string) {
	var t models.Product
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if t.ProdID == 0 {
		return 400, "Debe especificar el ID del Producto"
	}

	result, err2 := bd.SelectProduct(t)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar capturar el registro del producto " + strconv.Itoa(t.ProdID) + " > " + err.Error()
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
