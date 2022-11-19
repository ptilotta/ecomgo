package awsgo

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

// InicializoAWS es la función inicial para inicializar AWS
func InicializoAWS() {
	Ctx = context.TODO()
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))

	if err != nil {
		fmt.Println("Error al cargar la configurations .aws/config " + err.Error())
		fmt.Println("---------------------------------------------------------------------------------------------------------")
		panic("")
	}
}
