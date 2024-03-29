package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ptilotta/ecomgo/models"
)

func InsertOrder(o models.Orders) (int64, error) {
	fmt.Println("Comienza Registro")

	err := DbConnnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	/* Armo INSERT para el registro */
	sentencia := "INSERT INTO orders (Order_UserUUID, Order_Total, Order_AddId) VALUES ('"
	sentencia += o.Order_UserUUID + "'," + strconv.FormatFloat(o.Order_Total, 'f', -1, 64) + "," + strconv.Itoa(o.Order_AddId) + ")"

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

	// Grabo el detalle de Ordenes

	for _, od := range o.OrderDetails {
		sentencia = "INSERT INTO orders_detail (OD_OrderId, OD_ProdId, OD_Quantity, OD_Price) VALUES (" + strconv.Itoa(int(LastInsertId))
		sentencia = sentencia + "," + strconv.Itoa(od.OD_ProdId) + "," + strconv.Itoa(od.OD_Quantity) + ","
		sentencia = sentencia + strconv.FormatFloat(od.OD_Price, 'f', -1, 64) + ")"

		fmt.Println(sentencia)
		_, err = Db.Exec(sentencia)
		if err != nil {
			fmt.Println(err.Error())
			return 0, err
		}
	}

	fmt.Println("Insert Order > Ejecución exitosa ")
	return LastInsertId, nil
}

/*
func SelectOrder(o models.Orders) (models.Orders, error) {
	fmt.Println("Comienza SelectOrder")
	var Order models.Orders
	var sentencia string = "SELECT Order_Id, Order_UserUUID, Order_AddId, Order_Date, Order_Total, OD_Id, OD_ProdId, OD_Quantity, OD_Price "
	sentencia = sentencia + " FROM orders o JOIN orders_detail od ON o.Order_Id = od.OD_OrderId "
	sentencia = sentencia + " WHERE Order_Id = " + strconv.Itoa(o.Order_Id)

	err := DbConnnect()
	if err != nil {
		return Order, err
	}
	defer Db.Close()

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)
	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return Order, err
	}

	for rows.Next() {
		var OrderDate sql.NullTime
		var Order_AddId sql.NullInt32
		var OD_Id int64
		var OD_ProdId int64
		var OD_Quantity int64
		var OD_Price float64
		err := rows.Scan(&Order.Order_Id, &Order.Order_UserUUID, &Order_AddId, &OrderDate, &Order.Order_Total, &OD_Id, &OD_ProdId, &OD_Quantity, &OD_Price)
		Order.Order_Date = OrderDate.Time.String()
		Order.Order_AddId = int(Order_AddId.Int32)

		if err != nil {
			return Order, err
		}

		var od models.OrdersDetails
		od.OD_Id = int(OD_Id)
		od.OD_ProdId = int(OD_ProdId)
		od.OD_Quantity = int(OD_Quantity)
		od.OD_Price = OD_Price

		Order.OrderDetails = append(Order.OrderDetails, od)
	}

	fmt.Println("Select Order > Ejecución exitosa ")
	return Order, err
}
*/

func SelectOrders(user string, fechaDesde string, fechaHasta string, page int, orderId int) ([]models.Orders, error) {
	fmt.Println("Comienza SelectOrders")
	var Orders []models.Orders

	var sentencia string = "SELECT Order_Id, Order_UserUUID, Order_AddId, Order_Date, Order_Total FROM orders"

	if orderId > 0 {
		sentencia += " WHERE Order_Id = " + strconv.Itoa(orderId)
	} else {
		offset := 0
		if page == 0 {
			page = 1
		}
		if page > 1 {
			fmt.Println("page = " + strconv.Itoa(page))
			offset = (10 * (page - 1))
		}

		if len(fechaHasta) == 10 {
			fechaHasta = fechaHasta + " 23:59:59"
		}

		var where string
		var whereUser string = " Order_UserUUID = '" + user + "'"
		if len(fechaDesde) > 0 && len(fechaHasta) > 0 {
			where += " WHERE Order_Date BETWEEN '" + fechaDesde + "' AND '" + fechaHasta + "'"
		}
		if len(where) > 0 {
			where += " AND " + whereUser
		} else {
			where += " WHERE " + whereUser
		}

		limit := " LIMIT 10 "
		if offset > 0 {
			limit += " OFFSET " + strconv.Itoa(offset)
		}

		sentencia += where + limit
	}

	fmt.Println(sentencia)
	err := DbConnnect()
	if err != nil {
		return Orders, err
	}
	defer Db.Close()

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)
	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return Orders, err
	}

	for rows.Next() {
		var Order models.Orders
		var OrderDate sql.NullTime
		var OrderAddId sql.NullInt32
		err := rows.Scan(&Order.Order_Id, &Order.Order_UserUUID, &OrderAddId, &OrderDate, &Order.Order_Total)
		Order.Order_Date = OrderDate.Time.String()
		Order.Order_AddId = int(OrderAddId.Int32)

		if err != nil {
			return Orders, err
		}

		// Leo orders_detail por cada registro en 'orders'
		var rowsD *sql.Rows
		sentenciaD := "Select OD_Id, OD_ProdId, OD_Quantity, OD_Price FROM orders_detail WHERE OD_OrderId=" + strconv.Itoa(Order.Order_Id)
		fmt.Println(sentenciaD)
		rowsD, err = Db.Query(sentenciaD)

		for rowsD.Next() {

			var OD_Id int64
			var OD_ProdId int64
			var OD_Quantity int64
			var OD_Price float64

			err = rowsD.Scan(&OD_Id, &OD_ProdId, &OD_Quantity, &OD_Price)

			var od models.OrdersDetails
			od.OD_Id = int(OD_Id)
			od.OD_ProdId = int(OD_ProdId)
			od.OD_Quantity = int(OD_Quantity)
			od.OD_Price = OD_Price

			Order.OrderDetails = append(Order.OrderDetails, od)
		}
		Orders = append(Orders, Order)

		rowsD.Close()
	}

	fmt.Println("Select Orders > Ejecución exitosa ")
	return Orders, err
}
