package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var conn *gorm.DB
var err error

func Connection() {
	provider := os.Getenv("DB_PROVIDER")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	authDB := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, name)
	conn, err = gorm.Open(provider, authDB)
	if err != nil {
		panic(err)
	}
}

func AppConnection() {
	if err := godotenv.Load("../.env"); err != nil {
		if err := godotenv.Load("./.env"); err != nil {
			panic(err)
		}
	}
	Connection()
}
func GetDB() *gorm.DB {
	return conn
}
