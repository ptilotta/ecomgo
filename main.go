package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ptilotta/ecomgo/awsgo"
	"github.com/ptilotta/ecomgo/bd"
	"github.com/ptilotta/ecomgo/handlers"

	lambda "github.com/aws/aws-lambda-go/lambda"
)

type Respuesta struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {

	awsgo.InicializoAWS()

	if !validoParametros() {
		panic("Error en los parámetros. debe enviar 'SecretName', 'UserPoolID', 'Region'")
	}

	var res *events.APIGatewayProxyResponse
	path := request.RawPath
	method := request.RequestContext.HTTP.Method
	body := request.Body
	headers := request.Headers

	fmt.Println("----------------------------------------------------------------")
	fmt.Println("path = " + path)
	fmt.Println("method = " + method)
	fmt.Println("body = " + body)
	fmt.Println("----------------------------------------------------------------")

	bd.ReadSecret()

	status, message := handlers.Manejadores(path, method, body, headers)
	mensaje, _ := json.Marshal(&Respuesta{
		Status:  status,
		Message: message,
	})

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(mensaje),
	}
	return res, nil
}

func main() {

	lambda.Start(EjecutoLambda)
}

func validoParametros() bool {

	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro
	}
	_, traeParametro = os.LookupEnv("UserPoolID")
	if !traeParametro {
		return traeParametro
	}
	_, traeParametro = os.LookupEnv("Region")
	if !traeParametro {
		return traeParametro
	}

	return traeParametro
}
