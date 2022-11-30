package bd

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ptilotta/ecomgo/models"
	"github.com/ptilotta/ecomgo/secretm"
)

// Se asignará 1 sola vez el valor y se compartirá por todos los archivos del mismo package, para no tener que leer x veces el secreto
var SecretModel models.SecretRDSJson
var Db *sql.DB
var Email string
var Expirate int64

// Funcion central de conexión a la BD
func DbConnnect() error {

	var err error
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Conexión exitosa a la BD ")
	return nil
}

// UserExists chequea si un email ya se encuentra en la tabla users
func UserExists(email string) (error, bool) {
	fmt.Println("Comienza UserExists")

	err := DbConnnect()
	if err != nil {
		return err, false
	}
	defer Db.Close()

	/* Armo INSERT para el registro */
	sentencia := "SELECT 1 FROM users WHERE User_Email='" + email + "'"

	rows, err := Db.Query(sentencia)
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

// ProcesoToken proceso token para extraer sus valores
func ProcesoToken(tk string) (*models.Claim, bool, error) {
	miClave := []byte("")
	claims := &models.Claim{}

	fmt.Println("token = " + tk)

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado := UserExists(claims.Email)
		if encontrado == true {
			Email = claims.Email
			Expirate = claims.Expirate
		}
		return claims, encontrado, nil
	}
	if !tkn.Valid {
		return claims, false, errors.New("token Inválido")
	}
	return claims, false, err
}
