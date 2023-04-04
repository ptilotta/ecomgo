package bd

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ptilotta/ecomgo/models"
	"github.com/ptilotta/ecomgo/tools"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("Comienza Registro")

	err := DbConnnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	/* Armo INSERT para el registro */
	sentencia := "INSERT INTO category (Categ_Name, Categ_Path) VALUES ('" + c.CategName + "','" + c.CategPath + "')"

	var result sql.Result
	result, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}

	fmt.Println("Insert Category > Ejecución exitosa ")
	return LastInsertId, err
}

func UpdateCategory(c models.Category) error {
	fmt.Println("Comienza Update")

	err := DbConnnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	/* Armo UPDATE para el registro */
	sentencia := "UPDATE category SET "

	if len(c.CategName) > 0 {
		sentencia = sentencia + " Categ_Name = '" + tools.EscapeString(c.CategName) + "'"
	}

	if len(c.CategPath) > 0 {
		if !strings.HasSuffix(sentencia, "SET ") {
			sentencia = sentencia + ", "
		}
		sentencia = sentencia + "Categ_Path = '" + tools.EscapeString(c.CategPath) + "'"
	}

	sentencia = sentencia + " WHERE Categ_Id = " + strconv.Itoa(c.CategID)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Category > Ejecución exitosa ")
	return nil
}

func DeleteCategory(id int) error {
	fmt.Println("Comienza Delete")

	err := DbConnnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	/* Armo UPDATE para el registro */
	sentencia := "DELETE FROM category WHERE Categ_Id=" + strconv.Itoa(id)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Delete Category > Ejecución exitosa ")
	return nil
}

func SelectCategories(CategId int, Slug string) ([]models.Category, error) {
	fmt.Println("Comienza SelectCategories")
	var Categ []models.Category
	err := DbConnnect()
	if err != nil {
		return Categ, err
	}
	defer Db.Close()

	/* Armo SELECT para el registro */
	sentencia := "SELECT Categ_Id, Categ_Name, Categ_Path FROM category "
	if CategId > 0 {
		sentencia += "WHERE Categ_Id=" + strconv.Itoa(CategId)
	} else {
		if len(Slug) > 0 {
			sentencia += "WHERE Categ_Path LIKE '%" + strconv.Itoa(CategId) + "%'"
		}
	}

	fmt.Println(sentencia)
	var rows *sql.Rows
	rows, err = Db.Query(sentencia)
	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return Categ, err
	}

	rows, err = Db.Query(sentencia)

	for rows.Next() {
		var c models.Category
		var categId sql.NullInt32
		var categName sql.NullString
		var categPath sql.NullString

		err := rows.Scan(&categId, &categName, &categPath)

		c.CategID = int(categId.Int32)
		c.CategName = categName.String
		c.CategPath = categPath.String

		if err != nil {
			return Categ, err
		}

		Categ = append(Categ, c)
	}

	fmt.Println("Select Category > Ejecución exitosa ")
	return Categ, err
}
