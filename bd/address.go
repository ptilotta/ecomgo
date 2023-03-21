package bd

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ptilotta/ecomgo/models"
)

func InsertAddress(addr models.Address, User string) error {
	fmt.Println("Comienza Registro InsertAddress")

	err := DbConnnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	/* Armo INSERT para el registro */
	sentencia := "INSERT addresses (Add_UserId, Add_Address, Add_City, Add_State, Add_PostalCode, Add_Phone, Add_Title, Add_Name) VALUES ("
	sentencia += "'" + User + "','" + addr.AddAddress + "','" + addr.AddCity + "','" + addr.AddState + "','"
	sentencia += addr.AddPostalCode + "','" + addr.AddPhone + "','" + addr.AddTitle + "','" + addr.AddName + "')"

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("InsertAddress > Ejecución exitosa ")
	return nil
}

func UpdateAddress(addr models.Address) error {
	fmt.Println("Comienza UpdateAddress")

	err := DbConnnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "UPDATE addresses SET "
	if addr.AddAddress != "" {
		sentencia += "Add_Address='" + addr.AddAddress + "', "
	}
	if addr.AddCity != "" {
		sentencia += "Add_City='" + addr.AddCity + "', "
	}
	if addr.AddName != "" {
		sentencia += "Add_Name='" + addr.AddName + "', "
	}
	if addr.AddPhone != "" {
		sentencia += "Add_Phone='" + addr.AddPhone + "', "
	}
	if addr.AddPostalCode != "" {
		sentencia += "Add_PostalCode='" + addr.AddPostalCode + "', "
	}
	if addr.AddState != "" {
		sentencia += "Add_State='" + addr.AddState + "', "
	}
	if addr.AddTitle != "" {
		sentencia += "Add_Title='" + addr.AddTitle + "', "
	}
	sentencia, _ = strings.CutSuffix(sentencia, ", ")
	sentencia += " WHERE Add_Id = " + strconv.Itoa(addr.AddId)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("UpdateAddress > Ejecución exitosa ")
	return nil
}

// SelectAddreses Selecciona los datos de Addreses de un usuario en particular
func SelectAddreses(User string) ([]models.Address, error) {
	fmt.Println("Comienza SelectAddreses")
	Addr := []models.Address{}

	err := DbConnnect()
	if err != nil {
		return Addr, err
	}
	defer Db.Close()

	/* Armo SELECT */
	var sentencia string

	sentencia = "Select * FROM addresses WHERE Add_UserId='" + User + "'"

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)
	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return Addr, err
	}
	for rows.Next() {
		var a models.Address
		var addId sql.NullInt16
		var addAddress sql.NullString
		var addCity sql.NullString
		var addState sql.NullString
		var addPostalCode sql.NullString
		var addPhone sql.NullString
		var addTitle sql.NullString
		var addName sql.NullString

		err := rows.Scan(&addId, &addAddress, &addCity, &addState, &addPostalCode, &addPhone, &addTitle, &addName)
		if err != nil {
			return Addr, err
		}
		a.AddId = int(addId.Int16)
		a.AddAddress = addAddress.String
		a.AddCity = addCity.String
		a.AddState = addState.String
		a.AddPostalCode = addPostalCode.String
		a.AddPhone = addPhone.String
		a.AddTitle = addTitle.String
		a.AddName = addName.String
		Addr = append(Addr, a)
	}

	fmt.Println(Addr)
	fmt.Println("Select Addresses > Ejecución exitosa ")
	return Addr, nil
}

// DeleteUser borra el registro del usuario
func DeleteAddress(addr models.Address) error {
	fmt.Println("Comienza DeleteAddress")

	err := DbConnnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	/* Armo DELETE para el registro */
	sentencia := "DELETE FROM addresses WHERE Add_Id=" + strconv.Itoa(addr.AddId)

	_, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("DeleteUser > Ejecución exitosa ")
	return nil
}

func AddressExists(User string, AddId int) (error, bool) {
	fmt.Println("Comienza AddressExists")

	err := DbConnnect()
	if err != nil {
		return err, false
	}
	defer Db.Close()

	/* Armo INSERT para el registro */
	sentencia := "SELECT 1 FROM addresses WHERE Add_Id ='" + strconv.Itoa(AddId) + "' AND Add_UserId='" + User + "'"
	fmt.Println(sentencia)
	rows, err := Db.Query(sentencia)
	if err != nil {
		return err, false
	}

	var valor string
	rows.Next()
	rows.Scan(&valor)

	fmt.Println("AddressExists > Ejecución exitosa - valor devuelto " + valor)
	if valor == "1" {
		return nil, true
	} else {
		return nil, false
	}
}
