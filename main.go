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

	if validoParametros() == false {
		panic("Error en los parámetros. debe enviar 'SecretName', 'UserPoolID', 'Region'")
	}

	var res *events.APIGatewayProxyResponse
	path := request.RawPath
	method := request.RequestContext.HTTP.Method
	body := request.Body
	headers := request.Headers
	userPoolID := os.Getenv("UserPoolID")
	region := os.Getenv("Region")

	fmt.Println("----------------------------------------------------------------")
	fmt.Println("path = " + path)
	fmt.Println("method = " + method)
	fmt.Println("body = " + body)
	fmt.Println("----------------------------------------------------------------")

	bd.ReadSecret()
	status, message := handlers.Manejadores(path, method, body, headers, userPoolID, region)
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
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	if traeParametro == false {
		return false
	}

	_, traeParametro = os.LookupEnv("UserPoolID")
	if traeParametro == false {
		return false
	}

	_, traeParametro = os.LookupEnv("Region")
	if traeParametro == false {
		return false
	}

	return true
}
