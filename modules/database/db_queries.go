package database

import (
	ansi "DataTerm/modules/ANSI"
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
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

func ExecuteUserQuery(g *gocui.Gui, query string) error {
	output, err := g.View("Output")
	if err != nil {
		QueryError(output, err)
		log.Fatalln(err)
	}
	rows, err := CNX.Query(query)
	if err != nil {
		QueryError(output, err)
		return nil
	}
	columns, err := rows.Columns()
	if err != nil {
		QueryError(output, err)
		return nil
	}
	columnsType, err := rows.ColumnTypes()
	if err != nil {
		QueryError(output, err)
		return nil
	}

	output.Clear()
	fmt.Fprint(output, "| ")
	for i := 0; i < len(columns); i++ {
		fmt.Fprint(output, columns[i]+" | ")
	}
	fmt.Fprint(output, "\n| ")
	for i := 0; i < len(columns); i++ {
		fmt.Fprint(output, columnsType[i].DatabaseTypeName())
		fmt.Fprint(output, " | ")
	}

	return nil
}

func QueryError(output *gocui.View, err error) {
	output.Clear()
	fmt.Fprintf(output, ansi.ErrorColor, err)
}
