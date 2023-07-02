package database

import (
	ansi "DataTerm/modules/ANSI"
	"DataTerm/modules/credentials"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	CNX    *sql.DB
	err    error
	Creds  credentials.Credentials
	Driver string
)

func SetCredentials(creds credentials.Credentials, driver string) {
	Creds = creds
	Driver = driver
}

func OpenDB() {
	CNX, err = sql.Open(Driver, BuildConnectionString(Creds, Driver))
	if err != nil {
		log.Fatal("[DATABASE] Fail to open the database:" + err.Error())
	}
	err = CNX.Ping()
	if err != nil {
		log.Fatal("Fail to ping database:" + err.Error())
	} else {
		log.Printf(ansi.ValidColor, "[DATABASE CNX] Database connected (PONG!)")
	}
}

func BuildConnectionString(creds credentials.Credentials, driver string) string {
	templates := map[string]string{
		"mysql":    fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", creds.User, creds.Password, creds.Host, creds.Port, creds.Database),
		"postgres": fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", creds.Host, creds.Port, creds.User, creds.Password, creds.Database),
		// "odbc":     false,
	}
	return templates[driver]
}
