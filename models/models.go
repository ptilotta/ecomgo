package models

//SecretRDSJson es la estructura que devuelve Secret Manager
type SecretRDSJson struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

//Signup es la estructura que contiene los datos del registro
type SignUp struct {
	UserEmail     string `json:"UserEmail"`
	UserPassword  string `json:"UserPassword"`
	UserFirstName string `json:"UserFirstName"`
	UserLastName  string `json:"UserLastName"`
	UserBirdDate  string `json:"UserBirdDate,omitempty"`
}
