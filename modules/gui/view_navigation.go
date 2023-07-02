package gui

import (
	"DataTerm/modules/database"
	"fmt"

	"github.com/jroimartin/gocui"
)

var (
	currentTree        string = "tables"
	lastIndex          int
	selectedTableIndex int = 0
	tables             []string
	columns            []string
)

//Tables Tree

func navigationView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	//Tables view
	if v, err := g.SetView("Nav", 0, 3, maxX/3, maxY-1); err != nil {

		if err != gocui.ErrUnknownView {
			return err
		}
		if err := refreshTablesTreeView(g, v); err != nil {
			return nil
		}
	}
	return nil
}

func refreshTablesTreeView(g *gocui.Gui, v *gocui.View) error {
	//Tables view
	if _, err := g.SetCurrentView("Nav"); err != nil {
		return err
	}
	v.Clear()
	//View params
	v.Highlight = true
	v.SelBgColor = gocui.ColorCyan
	v.Title = "Tables"
	//Content
	displayTablesTree(v, tables)
	return nil
}

func displayTablesTree(v *gocui.View, tree []string) {
	v.Clear()
	_, heigth := v.Size()
	if selectedTableIndex+1 >= heigth {
		for i := selectedTableIndex + 1 - heigth; i < len(tree); i++ {
			if i == selectedTableIndex {
				fmt.Fprintln(v, "-", tree[i])
			} else {
				fmt.Fprintln(v, "|", tree[i])
			}
		}
	} else {
		for i := 0; i < len(tree); i++ {
			if i == selectedTableIndex {
				fmt.Fprintln(v, "-", tree[i])
			} else {
				fmt.Fprintln(v, "|", tree[i])
			}
		}
	}
}

//Table columns tree

func tableColumnTreeView(g *gocui.Gui, v *gocui.View) error {
	if _, err := g.SetCurrentView("Nav"); err != nil {
		return err
	}
	v.Clear()
	//View params
	v.Highlight = true
	v.SelBgColor = gocui.ColorCyan
	v.Title = "[Esc] " + tables[selectedTableIndex]
	//Content
	columns = database.GetAllColumns(tables[selectedTableIndex])
	displayTablesTree(v, columns)
	return nil
}
