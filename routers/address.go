package routers

import (
	"encoding/json"

	"github.com/ptilotta/ecomgo/bd"
	"github.com/ptilotta/ecomgo/models"
)

/*InsertAddress es la funcion para Insertar en la BD una Address de Usuario */
func InsertAddress(body string, User string) (int, string) {
	var t models.Address
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if t.AddAddress == "" {
		return 400, "Debe especificar el Address"
	}
	if t.AddName == "" {
		return 400, "Debe especificar el Name de la Address"
	}
	if t.AddTitle == "" {
		return 400, "Debe especificar el Title de la Address"
	}
	if t.AddCity == "" {
		return 400, "Debe especificar la City de la Address"
	}
	if t.AddPhone == "" {
		return 400, "Debe especificar el Phone de la Address"
	}
	if t.AddPostalCode == "" {
		return 400, "Debe especificar el PostalCode de la Address"
	}

	err = bd.InsertAddress(t, User)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el registro del Address para el ID de Usuario " + User + " > " + err.Error()
	}

	return 200, "InsertAddress OK"
}

/*UpdateAddress es la funcion para actualizar en la BD una Address de Usuario */
func UpdateAddress(body string, User string, id int) (int, string) {
	var t models.Address
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	t.AddId = id
	var encontrado bool
	err, encontrado = bd.AddressExists(User, t.AddId)
	if !encontrado {
		if err != nil {
			return 400, "Error al intentar buscar Address para el usuario " + User + " > " + err.Error()
		} else {
			return 400, "No se encuentra un registro de ID de Usuario asociado a esa ID de Address"
		}
	}

	err = bd.UpdateAddress(t)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar la actualización del Address para el ID de Usuario " + User + " > " + err.Error()
	}

	return 200, "UpdateAddress OK"
}

/*SelectAddress es la funcion para obtener la lista de las addresses de un usuario */
func SelectAddresses(User string) (int, string) {

	addr, err := bd.SelectAddreses(User)
	if err != nil {
		return 400, "Ocurrió un error al intentar obtener la lista de direcciones del usuario " + User + " > " + err.Error()
	}

	respJson, err := json.Marshal(addr)
	if err != nil {
		return 500, "Error al formatear los datos de las Addresses como JSON"
	}

	return 200, string(respJson)
}

/*DeleteAddress es la funcion para borrar una dirección de un usuario de la base */
func DeleteAddress(User string, id int) (int, string) {
	err, encontrado := bd.AddressExists(User, id)
	if !encontrado {
		if err != nil {
			return 400, "Error al intentar buscar Address para el usuario " + User + " > " + err.Error()
		} else {
			return 400, "No se encuentra un registro de ID de Usuario asociado a esa ID de Address"
		}
	}

	if encontrado == false {
		return 400, "No se encuentra un registro de ID de Usuario asociado a esa ID de Address"
	}

	err = bd.DeleteAddress(id)
	if err != nil {
		return 400, "Ocurrió un error al intentar borrar una dirección del usuario '" + User + "' > " + err.Error()
	}

	return 200, "Delete OK!"
}
