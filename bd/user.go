package bd

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

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

	if len(UFields.UserFirstName) == 0 && len(UFields.UserLastName) == 0 {
		errText := "UserFirstName or UserLastName are required"
		fmt.Println(errText)
		return errors.New(errText)
	}

	/* Armo UPDATE para el registro */
	sentencia := "UPDATE users SET "

	coma := ""
	if len(UFields.UserFirstName) > 0 {
		coma = ","
		sentencia += "User_FirstName='" + UFields.UserFirstName + "'"
	}
	if len(UFields.UserLastName) > 0 {
		sentencia += coma + "User_FirstName='" + UFields.UserLastName + "'"
	}
	sentencia += ", User_DateUpg='" + tools.FechaMySQL() + "' WHERE User_UUID='" + UFields.UserUUID + "'"

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update > Ejecución exitosa ")
	return nil
}

// SelectUser Selecciona los datos de un usuario en particular
func SelectUser(UserID string) (models.User, error) {
	fmt.Println("Comienza SelectUser")
	User := models.User{}

	err := DbConnnect()
	if err != nil {
		return User, err
	}
	defer Db.Close()

	/* Armo UPDATE para el registro */
	sentencia := "Select * FROM users WHERE User_UUID='" + UserID + "'"
	fmt.Println(sentencia)
	var rows *sql.Rows
	rows, err = Db.Query(sentencia)
	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return User, err
	}
	rows.Next()

	var firstName sql.NullString
	var lastName sql.NullString
	var dateUpg sql.NullTime
	rows.Scan(&User.UserUUID, &User.UserEmail, &firstName, &lastName, &User.UserStatus, &User.UserDateAdd, &dateUpg)

	User.UserFirstName = firstName.String
	User.UserLastName = lastName.String
	User.UserDateUpg = dateUpg.Time.String()

	fmt.Println("Select User > Ejecución exitosa ")
	return User, nil
}

// SelectUser Selecciona los datos de un usuario en particular
func SelectUsers(UFields models.ListUsers) ([]models.User, error) {
	fmt.Println("Comienza SelectUser")
	User := []models.User{}

	err := DbConnnect()
	if err != nil {
		return User, err
	}
	defer Db.Close()

	/* Armo SELECT */
	var offset int = (UFields.Page * 10) - 10
	var sentencia string

	if offset > 0 {
		sentencia = "Select * FROM users LIMIT 10 OFFSET " + strconv.Itoa(offset)
	} else {
		sentencia = "Select * FROM users LIMIT 10 OFFSET " + strconv.Itoa(offset)
	}

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)
	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return User, err
	}
	for rows.Next() {
		var u models.User
		var firstName sql.NullString
		var lastName sql.NullString
		var dateUpg sql.NullTime

		err := rows.Scan(&u.UserUUID, &u.UserEmail, &firstName, &lastName, &u.UserStatus, &u.UserDateAdd, &dateUpg)
		if err != nil {
			return User, err
		}
		u.UserFirstName = firstName.String
		u.UserLastName = lastName.String
		u.UserDateUpg = dateUpg.Time.String()
		User = append(User, u)
	}

	fmt.Println("Select User > Ejecución exitosa ")
	return User, nil
}

// DeleteUser borra el registro del usuario
func DeleteUser(UFields models.DeleteUser) error {
	fmt.Println("Comienza DeleteUser")

	err := DbConnnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	/* Armo DELETE para el registro */
	sentencia := "DELETE FROM users WHERE user_UUID='" + UFields.UserUUID + "'"

	_, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("DeleteUser > Ejecución exitosa ")
	return nil
}
