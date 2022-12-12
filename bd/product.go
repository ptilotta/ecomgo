package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ptilotta/ecomgo/models"
)

func InsertProduct(p models.Product) (int64, error) {
	fmt.Println("Comienza Registro")

	err := DbConnnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	/* Armo INSERT para el registro */
	sentencia := "INSERT INTO products (Prod_Title"

	if len(p.ProdDescription) > 0 {
		sentencia = sentencia + ", Prod_Description"
	}

	if p.ProdPrice > 0 {
		sentencia = sentencia + ", Prod_Price"
	}

	sentencia = sentencia + ") VALUES ('" + p.ProdTitle + "'"

	if len(p.ProdDescription) > 0 {
		sentencia = sentencia + ", '" + EscapeString(p.ProdDescription) + "'"
	}

	if p.ProdPrice > 0 {
		sentencia = sentencia + ", " + strconv.FormatFloat(p.ProdPrice, 'e', -1, 64)
	}

	sentencia = sentencia + ")"

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

	fmt.Println("Insert Product > Ejecución exitosa ")
	return LastInsertId, err
}
