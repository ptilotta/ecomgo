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

	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}

	awsgo.InicializoAWS()

	if validoParametros() == false {
		panic("Error en los parámetros. debe enviar 'SecretName'")
	}

	var res *events.APIGatewayProxyResponse
	path := request.RawPath
	method := ""
	body := request.Body

	fmt.Println("----------------------------------------------------------------")
	fmt.Println("path = " + path)
	fmt.Println("method = " + method)
	fmt.Println("body = " + body)
	fmt.Println("----------------------------------------------------------------")

	fmt.Println(request)
	fmt.Printf("Body Size = %d.\n", len(request.Body))
	fmt.Println("method = " + "")
	fmt.Println(request.Body)

	bd.ReadSecret()
	status, message := handlers.Manejadores(path, method, body)
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

	return true
}
