package Config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var DB *gorm.DB

// get from .env file
var server = goDotEnvVariable("MYSQLSERVER")
var port = goDotEnvVariable("MYSQLPORT")
var username = goDotEnvVariable("MYSQLUSER")
var password = goDotEnvVariable("MYSQLPASSWORD")
var database = goDotEnvVariable("MYSQLDATABASE")

func DbURL() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		server,
		port,
		database,
	)
}

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open(DbURL()), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	// database.AutoMigrate(&Book{})

	DB = database
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
