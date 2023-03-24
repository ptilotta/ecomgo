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

func SelectProduct(p models.Product, choice string, page int, pageSize int, orderType string, orderField string) (models.ProductResp, error) {
	fmt.Println("Comienza SelectProduct")
	var Resp models.ProductResp
	var Prod []models.Product

	err := DbConnnect()
	if err != nil {
		return Resp, err
	}
	defer Db.Close()

	/* Armo SELECT para el registro */
	var where, limit string
	var sentencia string = "SELECT Prod_Id, Prod_Title, Prod_Description, Prod_CreatedAt, Prod_Updated, Prod_Price, Prod_Status, Prod_Path, Prod_CategoryId, Prod_Stock FROM products "
	var sentenciaCount string = "SELECT count(*) as registros FROM products "

	switch choice {
	case "P":
		where = " WHERE Prod_Id = " + strconv.Itoa(p.ProdID)
	case "S":
		where = " WHERE UCASE(CONCAT(Prod_Title, Prod_Description)) LIKE '%" + strings.ToUpper(p.ProdSearch) + "%'"
	case "C":
		where = " WHERE Prod_CategoryId = " + strconv.Itoa(p.ProdCategId)
	}

	sentenciaCount += where

	var rows *sql.Rows
	rows, err = Db.Query(sentenciaCount)
	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return Resp, err
	}

	rows.Next()
	var regi sql.NullInt32
	err = rows.Scan(&regi)
	registros := int(regi.Int32)

	if page > 0 {
		if registros > pageSize {
			limit = " LIMIT " + strconv.Itoa(pageSize)
			if page > 1 {
				offset := pageSize * (page - 1)
				limit += " OFFSET " + strconv.Itoa(offset)
			}
		} else {
			limit = ""
		}
	}

	var orderBy string
	if len(orderField) > 0 {
		switch orderField {
		case "I":
			orderBy = " ORDER BY Prod_Id "
		case "T":
			orderBy = " ORDER BY Prod_Title "
		case "D":
			orderBy = " ORDER BY Prod_Description "
		case "F":
			orderBy = " ORDER BY Prod_CreatedAt "
		case "P":
			orderBy = " ORDER BY Prod_Price "
		case "S":
			orderBy = " ORDER BY Prod_Stock "
		case "C":
			orderBy = " ORDER BY Prod_CategoryId "
		}
		if orderType == "D" {
			orderBy += " DESC"
		}
	}

	sentencia += where + limit + orderBy

	fmt.Println(sentencia)

	rows, err = Db.Query(sentencia)

	for rows.Next() {
		var p models.Product
		var prodId sql.NullInt32
		var prodTitle sql.NullString
		var prodDescription sql.NullString
		var prodCreatedAt sql.NullTime
		var prodUpdated sql.NullTime
		var prodPrice sql.NullFloat64
		var prodStatus sql.NullInt16
		var prodPath sql.NullString
		var prodCategoryId sql.NullInt32
		var prodStock sql.NullInt32

		err := rows.Scan(&prodId, &prodTitle, &prodDescription, &prodCreatedAt, &prodUpdated, &prodPrice, &prodStatus, &prodPath, &prodCategoryId, &prodStock)

		if err != nil {
			return Resp, err
		}

		p.ProdID = int(prodId.Int32)
		p.ProdTitle = prodTitle.String
		p.ProdDescription = prodDescription.String
		p.ProdCreatedAt = prodCreatedAt.Time.String()
		p.ProdUpdated = prodUpdated.Time.String()
		p.ProdPrice = prodPrice.Float64
		p.ProdStatus = int(prodStatus.Int16)
		p.ProdPath = prodPath.String
		p.ProdCategId = int(prodCategoryId.Int32)
		p.ProdStock = int(prodStock.Int32)
		Prod = append(Prod, p)
	}

	Resp.CantRows = registros
	Resp.Rows = Prod

	fmt.Println("Select Product > Ejecución exitosa ")
	return Resp, err
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
