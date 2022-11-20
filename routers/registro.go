package routers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"

	"github.com/ptilotta/ecomgo/bd"
	"github.com/ptilotta/ecomgo/models"
)

/*Registro es la funcion para crear en la BD el registro de usuario */
func Registro(event events.APIGatewayProxyRequest) (int, string) {
	var t models.SignUp
	err := json.Unmarshal([]byte(event.Body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.UserEmail) == 0 {
		return 400, "El email de usuario es requerido "
	}

	if len(t.UserPassword) < 6 {
		return 400, "Debe especificar una contraseña de al menos 6 caracteres"
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

	err = bd.SignUp(t)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el registro de usuario " + err.Error()
	}

	return 200, "SignUp OK"
}
