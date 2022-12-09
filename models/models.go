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

type ListUsers struct {
	UserUUID       string `json:"userUUID"`
	UserUUID_Admin string `json:"userUUID_Admin"`
	Page           int    `json:"page"`
}

type DeleteUser struct {
	UserUUID_Admin string `json:"userUUID_Admin"`
	UserUUID       string `json:"userUUID"`
}
