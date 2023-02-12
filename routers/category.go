package routers

import (
	"encoding/json"
	"strconv"

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
func UpdateCategory(body string, User string) (int, string) {
	var t models.Category
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if t.CategID == 0 {
		return 400, "Debe especificar ID de la Categoría a actualizar"
	}

	if len(t.CategName) == 0 && len(t.CategPath) == 0 {
		return 400, "Debe especificar Categ_Name o Categ_Path para actualizar"
	}

	// Chequeamos que sea Admin quien hace la petición
	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	err2 := bd.UpdateCategory(t)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar el UPDATE de la categoria " + strconv.Itoa(t.CategID) + " > " + err.Error()
	}

	return 200, "Update OK"
}

/*SelectCategory es la funcion para leer el registro de categoría */
func SelectCategory(body string) (int, string) {
	var t models.Category
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if t.CategID == 0 {
		return 400, "Debe especificar el ID de la Categoría"
	}

	result, err2 := bd.SelectCategory(t)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar capturar el registro de la categoría " + strconv.Itoa(t.CategID) + " > " + err.Error()
	}

	Categ, err3 := json.Marshal(result)
	if err3 != nil {
		return 400, "Ocurrió un error al intentar convertir en JSON el registro de Categoría"
	}

	return 200, string(Categ)
}
