package bd

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ptilotta/ecomgo/models"
	"github.com/ptilotta/ecomgo/secretm"
	"github.com/ptilotta/ecomgo/tools"
)

func SignUp(signupFields models.SignUp) error {
	fmt.Println("Comienza Registro")

	// Capturo el Secreto y leo los valores de Secret Manager
	claves := secretm.GetSecret(os.Getenv("SecretName"))

	/* Abro la base con las credenciales de root */
	db, err := sql.Open("mysql", ConnStr(claves))
	if err != nil {
		fmt.Println("SignUp > " + err.Error())
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("SignUp > " + err.Error())
		return err
	}
	/*-------------------------------------------*/

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

	_, err = db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SignUp > Ejecución exitosa ")
	return nil
}
