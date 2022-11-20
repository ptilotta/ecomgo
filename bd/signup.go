package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ptilotta/ecomgo/models"
	"github.com/ptilotta/ecomgo/tools"
)

func SignUp(signupFields models.SignUp) error {
	fmt.Println("Comienza Registro")

	err := DbConnnect()
	if err != nil {
		return err
	}

	fmt.Println("SignUp > Conexión exitosa a la BD ")

	/* Armo INSERT para el registro */
	sentencia := "INSERT INTO users (User_Email, User_Password, User_FirstName, User_LastName, User_DateAdd"

	/* Campos Opcionales */

	if len(signupFields.UserBirdDate) > 0 {
		sentencia = sentencia + ", User_BirdDate"
	}

	/* ------------------ */

	sentencia = sentencia + ") VALUES ('" + signupFields.UserEmail + "','" + signupFields.UserPassword + "','" + signupFields.UserFirstName + "','" + signupFields.UserLastName + "','" + tools.FechaMySQL() + "'"

	/* Campos Opcionales */

	if len(signupFields.UserBirdDate) > 0 {
		sentencia = sentencia + ", '" + signupFields.UserBirdDate + "'"
	}

	/* ------------------ */

	sentencia = sentencia + ")"

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SignUp > Ejecución exitosa ")
	return nil
}
