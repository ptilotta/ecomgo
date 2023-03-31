package routers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ptilotta/ecomgo/bd"
	"github.com/ptilotta/ecomgo/models"
)

/*Registro es la funcion para crear en la BD el registro de usuario */
func UpdateUser(body string, User string) (int, string) {
	var t models.User
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.UserFirstName) == 0 {
		return 400, "Debe especificar el Nombre (FirstName) del Usuario"
	}

	if len(t.UserLastName) == 0 {
		return 400, "Debe especificar el Apellido (LastName) del Usuario"
	}

	_, encontrado := bd.UserExists(User)
	if encontrado == false {
		return 400, "No existe un usuario registrado con ese UUID"
	}

	err = bd.UpdateUser(t, User)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el registro del usuario " + User + " > " + err.Error()
	}

	return 200, "UpdateUser OK"
}

/*SelectUser es la funcion para obtener los datos de un usuario */
func SelectUser(body string, User string) (int, string) {
	_, encontrado := bd.UserExists(User)
	if encontrado == false {
		return 400, "No existe un usuario registrado con ese UUID"
	}

	row, err := bd.SelectUser(User)
	fmt.Println(row)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el registro del usuario " + User + " > " + err.Error()
	}

	respJson, err := json.Marshal(row)
	if err != nil {
		return 500, "Error al formatear los datos del usuario como JSON"
	}

	return 200, string(respJson)
}

/*SelectUsers es la funcion para obtener la lista de los usuarios en la base */
func SelectUsers(body string, User string, request events.APIGatewayV2HTTPRequest) (int, string) {
	// Proceso los parámetros recibidos
	var Page int
	if len(request.QueryStringParameters["page"]) == 0 {
		Page = 1
	} else {
		Page, _ = strconv.Atoi(request.QueryStringParameters["page"])
	}

	// Chequeamos que sea Admin quien hace la petición
	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	user, err := bd.SelectUsers(Page)
	if err != nil {
		return 400, "Ocurrió un error al intentar obtener la lista de usuarios > " + err.Error()
	}

	respJson, err := json.Marshal(user)
	if err != nil {
		return 500, "Error al formatear los datos de los usuarios como JSON"
	}

	return 200, string(respJson)
}

/*DeleteUser es la funcion para borrar un usuario de la base */
func DeleteUser(User string, id string) (int, string) {

	_, encontrado := bd.UserExists(id)
	if encontrado == false {
		return 400, "No existe un usuario registrado con ese UUID"
	}

	// Chequeamos que sea Admin quien hace la petición
	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	err := bd.DeleteUser(id)
	if err != nil {
		return 400, "Ocurrió un error al intentar borrar el usuario > " + err.Error()
	}

	return 200, "Delete OK!"
}
