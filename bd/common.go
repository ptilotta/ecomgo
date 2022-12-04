package bd

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	jwt "github.com/golang-jwt/jwt/v4"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ptilotta/ecomgo/models"
	"github.com/ptilotta/ecomgo/secretm"
)

// Se asignará 1 sola vez el valor y se compartirá por todos los archivos del mismo package, para no tener que leer x veces el secreto
var SecretModel models.SecretRDSJson
var Db *sql.DB
var UserName string
var Expirate int64

type UserClaim struct {
	jwt.RegistreredClaims
	UserName string `json:"username"`
}

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
func UserExists(userName string) (error, bool) {
	fmt.Println("Comienza UserExists")

	err := DbConnnect()
	if err != nil {
		return err, false
	}
	defer Db.Close()

	/* Armo INSERT para el registro */
	sentencia := "SELECT 1 FROM users WHERE User_UUID='" + userName + "'"

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
func ProcesoToken(tk string, userPoolID string, region string) (UserClaim, bool, string, error) {
	fmt.Println("================================================================================================")
	fmt.Println("Comienza ProcesoToken")
	fmt.Println("================================================================================================")
	miClave := "8Xi/PEzDz4P6m9cRMLGZ7ilcxBHIdZfnEgEpw/q4IwA="

	var userClaim UserClaim
	tkn, err := jwt.ParseWithClaims(tk, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(miClave), nil
	})
	if err == nil {
		fmt.Println("Irá a UserExists(claims.UserName)")
		fmt.Println(userClaim.UserName)
		_, encontrado := UserExists(userClaim.UserName)
		if encontrado == true {
			UserName = userClaim.UserName
		}
		return userClaim, encontrado, "", nil
	} else {
		fmt.Println(err.Error())
	}
	if !tkn.Valid {
		return userClaim, false, string(""), errors.New("token Inválido")
	}
	fmt.Println("================================================================================================")
	return userClaim, false, string(""), err
}
