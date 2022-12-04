package bd

import (
	"fmt"

	"github.com/ptilotta/ecomgo/models"
)

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
