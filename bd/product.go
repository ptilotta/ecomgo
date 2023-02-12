package bd

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ptilotta/ecomgo/models"
	"github.com/ptilotta/ecomgo/tools"
)

func InsertProduct(p models.Product) (int64, error) {
	fmt.Println("Comienza Registro")

	err := DbConnnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	/* Armo INSERT para el registro */
	sentencia := "INSERT INTO products (Prod_Title "

	if len(p.ProdDescription) > 0 {
		sentencia = sentencia + ", Prod_Description"
	}

	if p.ProdPrice > 0 {
		sentencia = sentencia + ", Prod_Price"
	}

	if p.ProdCategId > 0 {
		sentencia = sentencia + ", Prod_CategoryId"
	}

	if p.ProdStock > 0 {
		sentencia = sentencia + ", Prod_Stock"
	}

	if len(p.ProdPath) > 0 {
		sentencia = sentencia + ", Prod_Path"
	}

	sentencia = sentencia + ") VALUES ('" + tools.EscapeString(p.ProdTitle) + "'"

	if len(p.ProdDescription) > 0 {
		sentencia = sentencia + ", '" + tools.EscapeString(p.ProdDescription) + "'"
	}

	if p.ProdPrice > 0 {
		sentencia = sentencia + ", " + strconv.FormatFloat(p.ProdPrice, 'e', -1, 64)
	}

	if p.ProdCategId > 0 {
		sentencia = sentencia + ", " + strconv.Itoa(p.ProdCategId)
	}

	if p.ProdStock > 0 {
		sentencia = sentencia + ", " + strconv.Itoa(p.ProdStock)
	}

	if len(p.ProdPath) > 0 {
		sentencia = sentencia + ", '" + p.ProdPath + "'"
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

func SelectProduct(p models.Product) (models.Product, error) {
	fmt.Println("Comienza SelectProduct")
	var Prod models.Product
	err := DbConnnect()
	if err != nil {
		return Prod, err
	}
	defer Db.Close()

	/* Armo SELECT para el registro */
	sentencia := "SELECT Prod_Title, Prod_Description, Prod_CreatedAt, Prod_Updated, Prod_Price, Prod_Status, Prod_Path, Prod_CategoryId, Prod_Stock FROM products WHERE Prod_Id = " + strconv.Itoa(p.ProdID)

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)
	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return Prod, err
	}

	rows.Next()

	var prodTitle sql.NullString
	var prodDescription sql.NullString
	var prodCreatedAt sql.NullTime
	var prodUpdated sql.NullTime
	var prodPrice sql.NullFloat64
	var prodStatus sql.NullInt16
	var prodPath sql.NullString
	var prodCategoryId sql.NullInt32
	var prodStock sql.NullInt32
	rows.Scan(&prodTitle, &prodDescription, &prodCreatedAt, &prodUpdated, &prodPrice, &prodStatus, &prodPath, &prodCategoryId, &prodStock)

	Prod.ProdTitle = prodTitle.String
	Prod.ProdDescription = prodDescription.String
	Prod.ProdCreatedAt = prodCreatedAt.Time.String()
	Prod.ProdUpdated = prodUpdated.Time.String()
	Prod.ProdPrice = prodPrice.Float64
	Prod.ProdStatus = int(prodStatus.Int16)
	Prod.ProdPath = prodPath.String
	Prod.ProdCategId = int(prodCategoryId.Int32)
	Prod.ProdStock = int(prodStock.Int32)

	fmt.Println("Select Product > Ejecución exitosa ")
	return Prod, err
}

func UpdateProduct(p models.Product) error {
	fmt.Println("Comienza Update")

	err := DbConnnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	/* Armo UPDATE para el registro */
	sentencia := "UPDATE products SET "

	if len(p.ProdTitle) > 0 {
		sentencia = sentencia + " Prod_Title = '" + tools.EscapeString(p.ProdTitle) + "'"
	}

	if len(p.ProdDescription) > 0 {
		if !strings.HasSuffix(sentencia, "SET ") {
			sentencia = sentencia + ", "
		}
		sentencia = sentencia + "Prod_Description = '" + tools.EscapeString(p.ProdDescription) + "'"
	}

	if p.ProdPrice > 0 {
		if !strings.HasSuffix(sentencia, "SET ") {
			sentencia = sentencia + ", "
		}
		sentencia = sentencia + "Prod_Price = " + strconv.FormatFloat(p.ProdPrice, 'e', -1, 64)
	}

	if p.ProdCategId > 0 {
		if !strings.HasSuffix(sentencia, "SET ") {
			sentencia = sentencia + ", "
		}
		sentencia = sentencia + "Prod_CategoryId = " + strconv.Itoa(p.ProdCategId)
	}

	if p.ProdStock > 0 {
		if !strings.HasSuffix(sentencia, "SET ") {
			sentencia = sentencia + ", "
		}
		sentencia = sentencia + "Prod_Stock = " + strconv.Itoa(p.ProdStock)
	}

	if len(p.ProdPath) > 0 {
		if !strings.HasSuffix(sentencia, "SET ") {
			sentencia = sentencia + ", "
		}
		sentencia = sentencia + "Prod_Path = '" + p.ProdPath + "'"
	}

	sentencia = sentencia + " WHERE Prod_Id = " + strconv.Itoa(p.ProdID)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Product > Ejecución exitosa ")
	return nil
}

func DeleteProduct(p models.Product) error {
	fmt.Println("Comienza Delete")

	err := DbConnnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	/* Armo DELETE para el registro */
	sentencia := "DELETE FROM products WHERE Prod_Id = " + strconv.Itoa(p.ProdID)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Delete Product > Ejecución exitosa ")
	return nil
}

func UpdateStock(p models.Product) error {
	fmt.Println("Comienza Delete")

	if p.ProdStock == 0 {
		return errors.New("[ERROR] Debe enviar el Stock a modificar")
	}

	err := DbConnnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	/* Armo UPDATE para el Stock */
	sentencia := "UPDATE products SET Prod_Stock = Prod_Stock + " + strconv.Itoa(p.ProdStock) + " WHERE Prod_Id = " + strconv.Itoa(p.ProdID)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Stock > Ejecución exitosa ")
	return nil
}
