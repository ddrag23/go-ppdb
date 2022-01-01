package db

import (
	"database/sql"
	"ddrag23/ppdb/utils"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

var dbConn *sql.DB

func openConnection()  {
	// load env
	err := godotenv.Load(".env")
	utils.PanicIfError(err) 
	
	// get file sql
	stmt, err := ioutil.ReadFile("./db/sql/ppdb.sql")
	utils.PanicIfError(err)

	// connect db
	db,err := sql.Open("pgx",os.Getenv("DB_URL"))
	utils.PanicIfError(err)

	err = db.Ping()
	
	utils.PanicIfError(err)
	
	_,err = db.Exec(string(stmt))
	utils.PanicIfError(err)
	fmt.Println("Create db success")
	dbConn = db
}

func GetConnection() *sql.DB {
	return dbConn
}

func RunDB()  {
	openConnection()
}