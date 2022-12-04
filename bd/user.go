package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ptilotta/ecomgo/models"
	"github.com/ptilotta/ecomgo/tools"
)

func UpdateUser(UFields models.User) error {
	fmt.Println("Comienza Registro")

	err := DbConnnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	/* Armo UPDATE para el registro */
	sentencia := "UPDATE users SET User_FirstName='" + UFields.UserFirstName + "', User_LastName='" + UFields.UserLastName +
		"', User_DateUpg='" + tools.FechaMySQL() + "' WHERE User_UUID='" + UFields.UserUUID + "'"

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update > Ejecución exitosa ")
	return nil
}
