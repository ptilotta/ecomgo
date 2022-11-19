package bd

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ptilotta/ecomgo/models"
	"github.com/ptilotta/ecomgo/secretm"
)

// Se asignará 1 sola vez el valor y se compartirá por todos los archivos del mismo package, para no tener que leer x veces el secreto
var SecretModel models.SecretRDSJson

func UserExists(email string) (error, bool) {
	fmt.Println("Comienza UserExists")

	/* Abro la base con las credenciales de root */
	db, err := sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println("UserExists > " + err.Error())
		return err, false
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("UserExists > " + err.Error())
		return err, false
	}
	/*-------------------------------------------*/

	fmt.Println("UserExists > Conexión exitosa a la BD ")

	/* Armo INSERT para el registro */
	sentencia := "SELECT 1 FROM users WHERE User_Email='" + email + "'"

	rows, err := db.Query(sentencia)
	if err != nil {
		return err, false
	}

	var valor int
	rows.Scan(&valor)

	fmt.Println("UserExists > Ejecución exitosa - valor devuelto " + strconv.Itoa(valor))
	if valor == 1 {
		return nil, true
	} else {
		return nil, false
	}
}

// ReadSecret ejecuta GetSecret del módulo secretm y devuelve el modelo JSON a la variable Global SecretModel
func ReadSecret() {
	// Capturo el Secreto y leo los valores de Secret Manager
	SecretModel = secretm.GetSecret(os.Getenv("SecretName"))
}
