package main

import (
	"os"

	"github.com/ptilotta/ecomgo/awsgo"
	"github.com/ptilotta/ecomgo/bd"
	"github.com/ptilotta/ecomgo/handlers"

	lambda "github.com/aws/aws-lambda-go/lambda"
)

func EjecutoLambda() {
	awsgo.InicializoAWS()

	if validoParametros() == false {
		panic("Error en los parámetros. debe enviar 'SecretName'")
	}

	bd.ReadSecret()
	handlers.Manejadores()
}

func main() {
	lambda.Start(EjecutoLambda)
}

func validoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	if traeParametro == false {
		return false
	}

	return true
}
