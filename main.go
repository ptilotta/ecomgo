package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ptilotta/ecomgo/auth"
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
	token := "eyJraWQiOiJmMnNMNEt5a1AweFlBNDFieUp3TEQwTFlMOXZIVlltb0VpMFwvdmdqZ2ZGYz0iLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiI0Y2UyNTJmNS04ZmMwLTQxYjgtYTMwOS03NDc1YjhlMzc4NDMiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMV9sbjBiTnVSQnYiLCJ2ZXJzaW9uIjoyLCJjbGllbnRfaWQiOiI3aGw2MThubXZwbm9ha2Uzc3IyOTBkdHRhaCIsImV2ZW50X2lkIjoiNGNjOGFhZDUtZWNhYy00ZGUzLWJkMTItNDY3MjY0NmE1ZGJmIiwidG9rZW5fdXNlIjoiYWNjZXNzIiwic2NvcGUiOiJvcGVuaWQgZW1haWwiLCJhdXRoX3RpbWUiOjE2NzU1NzI1OTcsImV4cCI6MTY3NTY1ODk5NywiaWF0IjoxNjc1NTcyNTk3LCJqdGkiOiIwMzY5NDA3OC0yYTFjLTRhZTEtODIxYS1lZGUxMDA1NTU5MmQiLCJ1c2VybmFtZSI6IjRjZTI1MmY1LThmYzAtNDFiOC1hMzA5LTc0NzViOGUzNzg0MyJ9.XV05V2R91Ab-u5MOi-FWykglYQyyaKndJwlw28hU_fCnWvW0I9or2iEIf05NhiccFgebx6YVGEjCRi5ZI862tDltaNK3KXGmHx4Uw5ozujPc1tGFt4qvdX_x0qBnBIKUbQp6P_Y-2cH4AUg5uyWIL9-OCn2poIQN6R1t2miWtVtXTjB7T2CD1x4uVNya5gnkmFvld5cGiU5p7JNs1w6K4CZO4eu_kbR4q-2rQEqkRMxZJZH9hJtSZ4RGpN2TJvAsVrweanFz2C8F4X14mm3aDvp3gQ-BvwUVUbvqSYEyW-AhzII_2dXB6SBhaioHx-4PUEpUIpDUHeV01NQFJ4jBAw"
	todoOK, err2, msg := auth.ValidoToken(token)
	if !todoOK {
		if err2 != nil {
			fmt.Println("Error en el token " + err2.Error())
		} else {
			fmt.Println("Error en el token " + msg)
		}
	} else {
		fmt.Println("Token OK")
	}

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
