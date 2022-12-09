package routers

import (
	"encoding/json"

	"github.com/ptilotta/ecomgo/bd"
	"github.com/ptilotta/ecomgo/models"
)

/*Registro es la funcion para crear en la BD el registro de usuario */
func UpdateUser(body string) (int, string) {
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

	if len(t.UserUUID) == 0 {
		return 400, "Debe especificar el UUID del Usuario (dato username en Cognito)"
	}

	_, encontrado := bd.UserExists(t.UserUUID)
	if encontrado == false {
		return 400, "No existe un usuario registrado con ese UUID"
	}

	err = bd.UpdateUser(t)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el registro del usuario " + t.UserUUID + " > " + err.Error()
	}

	return 200, "UpdateUser OK"
}

/*SelectUser es la funcion para obtener los datos de un usuario */
func SelectUser(body string) (int, string) {
	var t models.User
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.UserUUID) == 0 {
		return 400, "Debe especificar el UUID del Usuario (dato username en Cognito)"
	}

	_, encontrado := bd.UserExists(t.UserUUID)
	if encontrado == false {
		return 400, "No existe un usuario registrado con ese UUID"
	}

	var user models.User
	user, err = bd.SelectUser(t)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el registro del usuario " + t.UserUUID + " > " + err.Error()
	}

	respJson, err := json.Marshal(user)
	if err != nil {
		return 500, "Error al formatear los datos del usuario como JSON"
	}

	return 200, string(respJson)
}

/*SelectUsers es la funcion para obtener la lista de los usuarios en la base */
func SelectUsers(body string) (int, string) {
	var t models.ListUsers
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if t.Page == 0 {
		t.Page = 1
	}

	_, encontrado := bd.UserIsAdmin(t.UserUUID_Admin)
	if encontrado == false {
		return 400, "No existe un usuario registrado con ese UUID, o no está configurado como Admin"
	}

	var user []models.User
	user, err = bd.SelectUsers(t)
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
func DeleteUsers(body string) (int, string) {
	var t models.DeleteUser
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.UserUUID) == 0 {
		return 400, "Debe especificar el UUID del Usuario (dato username en Cognito)"
	}

	if len(t.UserUUID_Admin) == 0 {
		return 400, "Debe especificar el UUID del Usuario Admin (dato username en Cognito)"
	}

	_, encontrado := bd.UserIsAdmin(t.UserUUID_Admin)
	if encontrado == false {
		return 400, "No existe un usuario registrado con ese UUID, o no está configurado como Admin"
	}

	_, encontrado = bd.UserExists(t.UserUUID)
	if encontrado == false {
		return 400, "No existe un usuario registrado con ese UUID"
	}

	err = bd.DeleteUser(t)
	if err != nil {
		return 400, "Ocurrió un error al intentar borrar el usuario > " + err.Error()
	}

	return 200, "Delete OK!"
}
