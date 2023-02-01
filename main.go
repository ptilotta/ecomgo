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
	"github.com/ptilotta/ecomgo/tools"

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
	ok, err, mesaje := tools.ValidoJWT("eyJraWQiOiJmMnNMNEt5a1AweFlBNDFieUp3TEQwTFlMOXZIVlltb0VpMFwvdmdqZ2ZGYz0iLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiI0Y2UyNTJmNS04ZmMwLTQxYjgtYTMwOS03NDc1YjhlMzc4NDMiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMV9sbjBiTnVSQnYiLCJ2ZXJzaW9uIjoyLCJjbGllbnRfaWQiOiI3aGw2MThubXZwbm9ha2Uzc3IyOTBkdHRhaCIsImV2ZW50X2lkIjoiM2RmYTA2MmItMWI1Ni00ZjgyLWEzNmUtYjgyY2QxN2ZhOWJkIiwidG9rZW5fdXNlIjoiYWNjZXNzIiwic2NvcGUiOiJvcGVuaWQgZW1haWwiLCJhdXRoX3RpbWUiOjE2NzA4NzczOTUsImV4cCI6MTY3MDk2Mzc5NSwiaWF0IjoxNjcwODc3Mzk1LCJqdGkiOiJmNzM3Njk3MS04ZDQwLTQ5MzMtOTBkMC0zOWUxYjM0ZjA4YmQiLCJ1c2VybmFtZSI6IjRjZTI1MmY1LThmYzAtNDFiOC1hMzA5LTc0NzViOGUzNzg0MyJ9.ERQ10hs63DmMOnmVKiEDlnit6EwHREzSZkgVaDjbBLzctbm2qmQLSRU7c8UOBemUvonnMfavqvdE1DNN8zaXE1STuOFwnMLjmndO3pmMJtOuXG9hllWmPUDD1u5MxQMVs0lG9FzwJIh122RryI8U0dz48Wi95GB4qkNiGjpvgNpxFvD54D5ZBq9fmi0oWXYp56ZuH-Os97rILPWRR8Mw7OFKKJ3vQlrCBrJnucCY9ZEuGDwoOxeFJ4tUxJxhpZEviEGax9b1NrNv7VYuEpztXEdw5jkl2tqY0JVvgDvW0uKGThvQZu7NTwzEwTdeRzBYo_vrn7T9UpdgILpj0ij9iQ")

	fmt.Println(mesaje)

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
