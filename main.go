package main

import (
	ansi "DataTerm/modules/ANSI"
	"DataTerm/modules/credentials"
	"DataTerm/modules/database"
	"DataTerm/modules/gui"
	"DataTerm/modules/input"
	"log"
)

func main() {
	//INIT
	ansi.PrintTitle()
	creds, driver := credentials.ReadFlags()
	log.Printf(ansi.ValidColor, "[INIT] init OK")
	//DATABASE
	creds.Password = input.GetPassword()
	database.SetCredentials(creds, driver)
	database.OpenDB()
	gui.Display()
}
