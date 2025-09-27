package initializers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var SQLDB *sql.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")

	//SQL DB
	SQLDB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Err connect SQL:", err)
	}
	if err = SQLDB.Ping(); err != nil {
		log.Fatal("Err ping SQL:", err)
	}
	fmt.Println("Connected SQL")

	//GORM DB
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected GORM")
}
