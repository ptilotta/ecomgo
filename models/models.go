package models

// SecretRDSJson es la estructura que devuelve Secret Manager
type SecretRDSJson struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

type User struct {
	UserUUID      string `json:"userUUID"`
	UserEmail     string `json:"userEmail"`
	UserFirstName string `json:"userFirstName"`
	UserLastName  string `json:"userLastName"`
	UserStatus    int    `json:"userStatus"`
	UserDateAdd   string `json:"userDateAdd"`
	UserDateUpg   string `json:"userDateUpg"`
}

type DeleteUser struct {
	UserUUID string `json:"userUUID"`
}

type ListUsers struct {
	TotalItems int    `json:"totalItems"`
	Data       []User `json:"data"`
}

type Product struct {
	ProdID          int     `json:"prodID"`
	ProdTitle       string  `json:"prodTitle"`
	ProdDescription string  `json:"prodDescription"`
	ProdCreatedAt   string  `json:"prodCreatedAt"`
	ProdUpdated     string  `json:"prodUpdated"`
	ProdPrice       float64 `json:"prodPrice,omitempty"`
	ProdStatus      int     `json:"prodStatus"`
	ProdStock       int     `json:"prodStock"`
	ProdCategId     int     `json:"prodCategId"`
	ProdPath        string  `json:"prodPath"`
	ProdSearch      string  `json:"search"`
}

type ProductResp struct {
	TotalItems int       `json:"totalItems"`
	Data       []Product `json:"data"`
}

type Category struct {
	CategID   int    `json:"categID"`
	CategName string `json:"categName"`
	CategPath string `json:"categPath"`
}

type OrdersDetails struct {
	OD_Id       int     `json:"odId"`
	OD_OrderId  int     `json:"odOrderId"`
	OD_ProdId   int     `json:"odProdId"`
	OD_Quantity int     `json:"odQuantity"`
	OD_Price    float64 `json:"odPrice"`
}

type Orders struct {
	Order_Id       int     `json:"orderId"`
	Order_UserUUID string  `json:"orderUserUUID"`
	Order_AddId    int     `json:"orderAddId"`
	Order_Date     string  `json:"orderDate"`
	Order_Total    float64 `json:"orderTotal"`
	OrderDetails   []OrdersDetails
}

type Address struct {
	AddId         int    `json:"addId"`
	AddTitle      string `json:"addTitle"`
	AddName       string `json:"addName"`
	AddAddress    string `json:"addAddress"`
	AddCity       string `json:"addCity"`
	AddState      string `json:"addState"`
	AddPostalCode string `json:"addPostalCode"`
	AddPhone      string `json:"addPhone"`
}
