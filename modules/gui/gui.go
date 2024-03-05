package gui

import (
	"DataTerm/modules/database"
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

// INIT
func Display() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.InputEsc = true
	g.Highlight = true
	if err := layout(g); err != nil {
		log.Fatalln(err)
	}
	initKeyBinding(g)
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalln(err)
	}
}

// DEF
func layout(g *gocui.Gui) error {
	//Tables view
	tables = database.GetAllTables()
	g.SetManagerFunc(func(g *gocui.Gui) error {
		if err := navigationView(g); err != nil {
			return err
		}
		if err := headerView(g); err != nil {
			return err
		}
		if err := testView(g); err != nil {
			return err
		}
		//TODO OUTPUT VIEW
		if err := outputView(g); err != nil {
			return err
		}
		//TODO CONSOLE VIEW
		if err := consoleView(g); err != nil {
			return err
		}
		return nil
	})

	return nil
}

func testView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	//Title view
	if v, err := g.SetView("test", maxX/2, maxY/2, maxX-2, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		refreshTest(g, v)
	}
	return nil
}

func refreshTest(g *gocui.Gui, v *gocui.View) error {
	test, _ := g.View("test")
	test.Clear()
	fmt.Fprintf(test, "Line selected = %d ", selectedTableIndex)
	fmt.Fprintf(test, "Total size = %d ", len(tables))
	return nil
}
