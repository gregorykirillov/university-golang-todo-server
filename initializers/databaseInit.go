package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	connection_string := LoadEnvVariables("DB_CONN_STRING")
	DB, err = gorm.Open(postgres.Open(connection_string), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to db")
	}
	fmt.Println("Established a successful connection with database")

}
