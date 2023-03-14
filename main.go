package main

import (
	"context"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ptilotta/ecomgo/awsgo"
	"github.com/ptilotta/ecomgo/bd"
	"github.com/ptilotta/ecomgo/handlers"

	lambda "github.com/aws/aws-lambda-go/lambda"
)

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

	bd.ReadSecret()

	status, message := handlers.Manejadores(path, method, body, headers, request)
	//mensaje, _ := json.Marshal(message)

	headersResp := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		//		StatusCode: status,
		StatusCode: http.StatusNoContent,
		Body:       string(message),
		Headers:    headersResp,
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
