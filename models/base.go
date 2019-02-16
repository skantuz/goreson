package models

import (
	"fmt"
	"os"

	"github.com/skantuz/godotenv"
	"github.com/skantuz/gorm"
	_ "github.com/skantuz/gorm/dialects/mssql"
	_ "github.com/skantuz/gorm/dialects/mysql"
	_ "github.com/skantuz/gorm/dialects/postgres"
	_ "github.com/skantuz/gorm/dialects/sqlite"
)

var db *gorm.DB //database

func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	dbUser := os.Getenv("db_user")
	dbPass := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbType := os.Getenv("db_type")
	dbUri := ""
	switch dbType {
	case "postgres":
		dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", dbHost, dbPort, dbUser, dbName, dbPass) //Build connection string Postgres
	case "mysql":
		dbUri := fmt.Sprintf("%s:%s@%s:%s/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName) //Build connection string Mysql
	case "mssql":
		dbUri := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, dbPass, dbHost, dbPort, dbName) //Build connection string Microsoft SQL
	case "sqlite3":
		dbUri := fmt.Sprintf("%s", dbName) //Build connection string Sqlite3
	}

	conn, err := gorm.Open(dbType, dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&User{}) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
