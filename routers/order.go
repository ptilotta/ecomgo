package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ptilotta/ecomgo/bd"
	"github.com/ptilotta/ecomgo/models"
)

/*SelectOrder es la funcion para leer el registro de orders y order_details */
func SelectOrder(body string, request events.APIGatewayV2HTTPRequest) (int, string) {
	var o models.Orders
	var err error

	// Proceso los parámetros recibidos
	if len(request.QueryStringParameters["orderId"]) == 0 {
		return 400, "Debe especificar el ID de la Categoría"
	} else {
		o.Order_Id, err = strconv.Atoi(request.QueryStringParameters["orderId"])
	}

	result, err2 := bd.SelectOrder(o)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar capturar el registro de la orden " + strconv.Itoa(o.Order_Id) + " > " + err.Error()
	}

	if result.Order_Total == 0 {
		return 400, "Orden " + strconv.Itoa(o.Order_Id) + " no existe "
	}

	Order, err3 := json.Marshal(result)
	if err3 != nil {
		return 400, "Ocurrió un error al intentar convertir en JSON el registro de Orden"
	}

	return 200, string(Order)
}

/*InsertOrder es la funcion para crear en la BD el registro de orden */
func InsertOrder(body string, User string) (int, string) {
	var o models.Orders
	err := json.Unmarshal([]byte(body), &o)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	OK, message := ValidOrder(o)
	if !OK {
		return 400, message
	}

	// Chequeamos que sea Admin quien hace la petición
	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertOrder(o)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar el registro de la orden > " + err.Error()
	}

	return 200, "{ OrderID: " + strconv.Itoa(int(result)) + "}"
}

func ValidOrder(o models.Orders) (bool, string) {
	if o.Order_Total == 0 {
		return false, "Debe indicar el total de la orden"
	}
	if len(o.Order_UserUUID) == 0 {
		return false, "Debe indicar el Usuario que ha comprado en la orden"
	}
	count := 0
	for _, od := range o.OrderDetails {
		if od.OD_ProdId == 0 {
			return false, "Debe indicar el ID del producto en el detalle en la orden"
		}
		if od.OD_Quantity == 0 {
			return false, "Debe indicar la cantidad del producto vendido en el detalle en la orden"
		}
		count++
	}
	if count == 0 {
		return false, "Debe indicar items en la orden"
	}
	return true, ""
}

/*SelectOrders es la funcion para leer las ordenes por rango de fechas y paginado */
func SelectOrders(user string, request events.APIGatewayV2HTTPRequest) (int, string) {
	var err error
	var fechaDesde, fechaHasta string
	var page int

	// Proceso los parámetros recibidos
	if len(request.QueryStringParameters["fechaDesde"]) > 0 {
		fechaDesde = request.QueryStringParameters["fechaDesde"]
	}
	if len(request.QueryStringParameters["fechaHasta"]) > 0 {
		fechaHasta = request.QueryStringParameters["fechaHasta"]
	}
	if len(request.QueryStringParameters["page"]) > 0 {
		page, _ = strconv.Atoi(request.QueryStringParameters["fechaHasta"])
	}

	result, err2 := bd.SelectOrders(fechaDesde, fechaHasta, page)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar capturar los registros ordenes del " + fechaDesde + " al " + fechaHasta + " > " + err.Error()
	}

	if result[0].Order_Total == 0 {
		return 400, "No hay Ordenes para ese rango de fechas "
	}

	Orders, err3 := json.Marshal(result)
	if err3 != nil {
		return 400, "Ocurrió un error al intentar convertir en JSON el registro de Orden"
	}

	return 200, string(Orders)
}
