package gui

import (
	"DataTerm/modules/database"

	"github.com/jroimartin/gocui"
)

func kb_Console(g *gocui.Gui) error {
	if err := g.SetKeybinding("Console", gocui.KeyEnter, gocui.ModNone, executeQuery); err != nil {
		return err
	}
	if err := g.SetKeybinding("Console", gocui.KeyCtrlL, gocui.ModNone, clearConsole); err != nil {
		return err
	}
	return nil
}

func executeQuery(g *gocui.Gui, v *gocui.View) error {
	console, err := g.View("Console")
	if err != nil {
		return err
	}
	database.ExecuteUserQuery(g, console.Buffer())
	return nil
}

func clearConsole(g *gocui.Gui, v *gocui.View) error {
	console, err := g.View("Console")
	if err != nil {
		return err
	}
	console.Clear()
	return nil
}
