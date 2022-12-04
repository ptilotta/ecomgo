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

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type UserClaim struct {
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

// ConnStr arma el String de conexión de la base de datos
func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string

	dbUser = claves.Username
	dbEndpoint = fmt.Sprintf("%s:%d", claves.Host, claves.Port)
	dbName = "ecommerce"
	authToken = claves.Password
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
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
func ProcesoToken(tk string, userPoolID string, region string) (string, bool, string, error) {
	fmt.Println("================================================================================================")
	fmt.Println("Comienza ProcesoToken")
	fmt.Println("================================================================================================")
	miClave := "8Xi/PEzDz4P6m9cRMLGZ7ilcxBHIdZfnEgEpw/q4IwA="

	tkn, err := jwt.ParseWithClaims(tk, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(miClave), nil
	})

	fmt.Println("Paso por aca 1")
	var UserUUID string

	if claims, ok := tkn.Claims.(*UserClaim); ok && tkn.Valid {
		fmt.Printf("%v %v", claims.UserName, claims.RegisteredClaims.Issuer)
		UserUUID = claims.UserName
	} else {
		fmt.Println(err)
		return "", false, err.Error(), err
	}

	fmt.Println("Paso por aca 2")

	if err == nil {
		fmt.Println("Irá a UserExists(claims.UserName)")
		fmt.Println(UserUUID)
		_, encontrado := UserExists(UserUUID)
		if encontrado == true {
			UserName = UserUUID
		}
		return UserUUID, encontrado, "", nil
	} else {
		fmt.Println(err.Error())
	}
	if !tkn.Valid {
		return UserUUID, false, string(""), errors.New("token Inválido")
	}
	fmt.Println("================================================================================================")
	return UserUUID, false, string(""), err
}
