package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

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

	token := "eyJraWQiOiJlS3lvdytnb1wvXC9yWmtkbGFhRFNOM25jTTREd0xTdFhibks4TTB5b211aE09IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiJjMTcxOGY3OC00ODY5LTRmMmEtYTk2ZS1lYmEwYmJkY2RkMjEiLCJldmVudF9pZCI6IjZmYWMyZGNjLTJlMzUtMTFlOS05NDZjLTZiZDI0YmRlZjFiNiIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE1NDk5MTQyNjUsImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC51cy13ZXN0LTIuYW1hem9uYXdzLmNvbVwvdXMtd2VzdC0yX0wwVldGSEVueSIsImV4cCI6MTU0OTkxNzg2NSwiaWF0IjoxNTQ5OTE0MjY1LCJqdGkiOiIzMTg0MDdkMC0zZDNhLTQ0NDItOTMyYy1lY2I0MjQ2MzRiYjIiLCJjbGllbnRfaWQiOiI2ZjFzcGI2MzZwdG4wNzRvbjBwZGpnbms4bCIsInVzZXJuYW1lIjoiYzE3MThmNzgtNDg2OS00ZjJhLWE5NmUtZWJhMGJiZGNkZDIxIn0.rJl9mdCrw_lertWhC5RiJcfhRP-xwTYkPLPXmi_NQEO-LtIJ-kwVEvUaZsPnBXku3bWBM3V35jdJloiXclbffl4SDLVkkvU9vzXDETAMaZEzOY1gDVcg4YzNNR4H5kHnl-G-XiN5MajgaWbjohDHTvbPnqgW7e_4qNVXueZv2qfQ8hZ_VcyniNxMGaui-C0_YuR6jdH-T14Wl59Cyf-UFEyli1NZFlmpUQ8QODGMUI12PVFOZiHJIOZ3CQM_Xs-TlRy53RlKGFzf6RQfRm57rJw_zLyJHHnB8DZgbdCRfhNsqZka7ZZUUAlS9aMzdmSc3pPFSJ-hH3p8eFAgB4E71g"

	// Separamos el token en tres partes, separadas por "."
	parts := strings.Split(token, ".")

	// Validamos que tengan 3 partes
	if len(parts) != 3 {
		fmt.Println("El token no es válido.")
	}

	// La segunda parte contiene la información de usuario codificada
	// como una estructura JSON. Debemos decodificarla.
	// En este ejemplo, usamos base64.StdEncoding.DecodeString para
	// decodificar la parte en una cadena.
	userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("No se puede decodificar la parte del token:", err)
	}
	// Aquí puedes hacer algo con la información de usuario, por ejemplo,
	// deserializarla en una estructura de datos
	fmt.Println("Información de usuario:", string(userInfo))

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
