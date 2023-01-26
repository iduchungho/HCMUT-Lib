package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"example/hcmut-lib/pkg/model"
)
// type Database struct {
// 	DB *gorm.DB
// }
var DB *gorm.DB

func creteSchema(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Document{})
}

func  DatabaseConection() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("failed to load env file")
	}

	// Load env properties
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// mysql", "root:password@tcp(localhost:3306)/testdb
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	// fmt.Print(dsn);

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	// create the schema
	creteSchema(db)

	DB = db
	fmt.Println("Database Connected")

}
