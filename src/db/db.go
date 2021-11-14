package db

import (
	"database/sql"
	"fmt"
	"time"

	f "github.com/lucas/hackathon/src/functions"

	//driver para conectar no banco de dados da ORACLE
	_ "github.com/go-sql-driver/mysql"
)

//Pool of connections
var Pool *sql.DB

func init() {
	fmt.Println("Carregando Banco de dados")
	connect()
}

//função apra conectar no banco de dados
func connect() {
	db, err := sql.Open("mysql", f.GoDotEnvVariable("userDB")+":"+f.GoDotEnvVariable("passwordDB")+
		"@tcp("+f.GoDotEnvVariable("hostDB")+":"+f.GoDotEnvVariable("portDB")+")/"+f.GoDotEnvVariable("nameDB"))
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	Pool = db
	fmt.Println("Banco de dados iniciado")
}
