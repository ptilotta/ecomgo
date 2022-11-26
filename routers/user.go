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

	_, encontrado := bd.UserExists(t.UserEmail)
	if encontrado == true {
		return 400, "Ya existe un usuario registrado con ese email"
	}

	err = bd.UpdateUser(t)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el registro de usuario " + err.Error()
	}

	return 200, "SignUp OK"
}
