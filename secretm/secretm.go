package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/ptilotta/ecomgo/awsgo"
	"github.com/ptilotta/ecomgo/models"
)

// GetSecret es la función que devuelve la password de Secret Manager
func GetSecret(nombreSecret string) (datos models.SecretRDSJson) {
	var datosSecret models.SecretRDSJson
	fmt.Println(" > Pido Secreto : " + nombreSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})
	if err != nil {
		fmt.Println(err.Error())
		panic("")
	}
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(" > Lectura Secreto OK : " + nombreSecret)
	return datosSecret
}
