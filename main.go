package main

import (
	ansi "DataTerm/modules/ANSI"
	"DataTerm/modules/credentials"
	"DataTerm/modules/database"
	"DataTerm/modules/gui"
	"DataTerm/modules/input"
	"fmt"
)

func main() {
	//INIT
	ansi.PrintTitle()
	creds, driver := credentials.ReadFlags()
	//DATABASE
	creds.Password = input.GetPassword()
	database.SetCredentials(creds, driver)
	database.OpenDB()
	gui.Display()
	fmt.Print("\033[H\033[2J")
}
