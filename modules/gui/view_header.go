package gui

import (
	ansi "DataTerm/modules/ANSI"
	"DataTerm/modules/database"
	"fmt"

	"github.com/jroimartin/gocui"
)

func headerView(g *gocui.Gui) error {
	maxX, _ := g.Size()
	//Title view
	if v, err := g.SetView("Title", 0, 0, maxX, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintf(v, fmt.Sprintf(ansi.WarningColor, "DataTerm-- "))
		fmt.Fprintf(v, "connected to ")
		fmt.Fprintf(v, fmt.Sprintf(ansi.Green, database.Creds.User+"@"+database.Creds.Host))
		fmt.Fprintf(v, " on ")
		fmt.Fprintf(v, fmt.Sprintf(ansi.B_Cyan, database.Creds.Database))
		fmt.Fprintf(v, " database")
		fmt.Fprintf(v, "%s", len(tables))
	}

	//Made by view
	if v, err := g.SetView("madeBy", maxX-maxX/6, 0, maxX-1, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintf(v, "> made by")
		fmt.Fprintf(v, fmt.Sprintf(ansi.Purple, " I3UX_"))
	}
	return nil
}
