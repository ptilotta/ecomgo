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
	UserUUID      string `json:"user_uuid"`
	UserEmail     string `json:"user_email"`
	UserFirstName string `json:"user_firstname"`
	UserLastName  string `json:"user_lastname"`
	UserStatus    int    `json:"user_status"`
	UserDateAdd   string `json:"user_dateadd"`
	UserDateUpg   string `json:"user_dateupg"`
	Page          int    `json:"page"`
}

type DeleteUser struct {
	UserUUID_Admin string `json:"user_uuid_admin"`
	UserUUID       string `json:"user_uuid"`
}
