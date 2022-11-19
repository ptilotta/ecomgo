package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ptilotta/ecomgo/bd"
	"github.com/ptilotta/ecomgo/models"
)

/*Registro es la funcion para crear en la BD el registro de usuario */
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.SignUp
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.UserEmail) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}
	if len(t.UserPassword) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}
	if len(t.UserFirstName) == 0 {
		http.Error(w, "Debe especificar el Nombre (FirstName) del Usuario", 400)
		return
	}
	if len(t.UserLastName) == 0 {
		http.Error(w, "Debe especificar el Apellido (LastName) del Usuario", 400)
		return
	}

	_, encontrado := bd.UserExists(t.UserEmail)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	err = bd.SignUp(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
