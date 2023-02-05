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
	token := "eyJraWQiOiI4WGlcL1BFekR6NFA2bTljUk1MR1o3aWxjeEJISWRaZm5FZ0Vwd1wvcTRJd0E9IiwiYWxnIjoiUlMyNTYifQ.eyJhdF9oYXNoIjoicVcwYS1CcGNXT0x1Z3hqdWRBdGlRQSIsInN1YiI6IjRjZTI1MmY1LThmYzAtNDFiOC1hMzA5LTc0NzViOGUzNzg0MyIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMV9sbjBiTnVSQnYiLCJjb2duaXRvOnVzZXJuYW1lIjoiNGNlMjUyZjUtOGZjMC00MWI4LWEzMDktNzQ3NWI4ZTM3ODQzIiwiYXVkIjoiN2hsNjE4bm12cG5vYWtlM3NyMjkwZHR0YWgiLCJldmVudF9pZCI6IjRjYzhhYWQ1LWVjYWMtNGRlMy1iZDEyLTQ2NzI2NDZhNWRiZiIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNjc1NTcyNTk3LCJleHAiOjE2NzU2NTg5OTcsImlhdCI6MTY3NTU3MjU5NywianRpIjoiZDJkYjFjOWMtYTEwYy00ZGUzLWIzNjItMDM4ZTAzYjk4NWI4IiwiZW1haWwiOiJ0aWxvdHRhcGFibG8rMjdAZ21haWwuY29tIn0.BbCI9da1GX0ZuamxBb56uip3u38EUGK1tx5_O_sB3plvapH7_CTFLgBxephLuZbxz4sHYSxUb8j65UvVwxG1MyzKTJWXxgED_h1uyi5t2m6g1HbWDVG-_ELuqquGY9ESwocTCuvDQxpN5ptC_dNeodAsVNPrB-T-IFp9vQ-8ESkqxKK2P2FareddBMKDrfKb3n5bRGWG8efLhXUsNOIyp_-OiDBXKxlDA5R4WcNTEEeVQMxy0z3wyF56jVoVRkUl7nQJlpMyZa1wXFbPhGr1Bxf1XVmkhlKeaEggYWhjVgZqWroVMt3-q7oJnXLTYE0o2atoCAXp1VsMKij1lqO1MA"
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
