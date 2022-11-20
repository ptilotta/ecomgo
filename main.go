package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ptilotta/ecomgo/awsgo"
	"github.com/ptilotta/ecomgo/bd"
	"github.com/ptilotta/ecomgo/handlers"

	lambda "github.com/aws/aws-lambda-go/lambda"
)

func EjecutoLambda(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println(request)
	fmt.Printf("Body Size = %d.\n", len(request.Body))
	fmt.Println(request.Body)
	awsgo.InicializoAWS()

	if validoParametros() == false {
		panic("Error en los parámetros. debe enviar 'SecretName'")
	}

	bd.ReadSecret()
	status, message := handlers.Manejadores(request)
	return respuesta(message, status), nil
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

func respuesta(message string, status int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: message, StatusCode: status}
}
