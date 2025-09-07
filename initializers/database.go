package initializers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//mysql:"//root:password@localhost:3306/meal-order-db"
	//dsn := "user:password@tcp(127.0.0.1:3306)/meal-order-db?charset=utf8mb4&parseTime=True&loc=Local"

	//----> Load mysql database url from env file.
	dsn := os.Getenv("DATABASE_URL")

	//----> Connect gorm to mysql database
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//----> Check for error.
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

}
