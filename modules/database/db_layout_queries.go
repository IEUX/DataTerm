package database

import (
	ansi "DataTerm/modules/ANSI"
	"log"
)

func GetAllTables() []string {
	var table string
	var tablesNames []string
	var query string = "SELECT table_name FROM information_schema.tables WHERE table_type='BASE TABLE' AND table_schema =?;"

	rows, err := CNX.Query(query, Creds.Database)
	if err != nil {
		log.Fatalf(ansi.ErrorColor, err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&table)
		if err != nil {
			log.Fatal(err)
		}
		tablesNames = append(tablesNames, table)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return tablesNames
}

func GetAllColumns(table string) []string {
	var (
		column  string
		columns []string
		query   string = "SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = ? ORDER BY ORDINAL_POSITION;"
	)
	rows, err := CNX.Query(query, table)
	if err != nil {
		log.Fatalf(ansi.ErrorColor, err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&column)
		if err != nil {
			log.Fatal(err)
		}
		columns = append(columns, column)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return columns
}
