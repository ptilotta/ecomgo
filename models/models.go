package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// SecretRDSJson es la estructura que devuelve Secret Manager
type SecretRDSJson struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

// Signup es la estructura que contiene los datos del registro
type User struct {
	UserEmail     string `json:"UserEmail"`
	UserFirstName string `json:"UserFirstName"`
	UserLastName  string `json:"UserLastName"`
	UserUUID      string `json:"UserUUID"`
}

/*Claim es la estructura usada para procesar el JWT*/
type Claim struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}
