package gui

import (
	"github.com/jroimartin/gocui"
)

func kb_Navigation(g *gocui.Gui) error {
	if err := g.SetKeybinding("Nav", gocui.KeyArrowUp, gocui.ModNone, previousTable); err != nil {
		return err
	}
	if err := g.SetKeybinding("Nav", gocui.KeyArrowDown, gocui.ModNone, nextTable); err != nil {
		return err
	}
	if err := g.SetKeybinding("Nav", gocui.KeyEnter, gocui.ModNone, enter); err != nil {
		return err
	}
	if err := g.SetKeybinding("Nav", gocui.KeyEsc, gocui.ModNone, escape); err != nil {
		return err
	}
	if err := g.SetKeybinding("Nav", gocui.KeyCtrlT, gocui.ModNone, refreshTest); err != nil {
		return err
	}
	return nil
}

func previousTable(g *gocui.Gui, v *gocui.View) error {
	var tree []string
	if currentTree == "tables" {
		tree = tables
	}
	if currentTree == "columns" {
		tree = columns
	}
	if selectedTableIndex > 0 {
		selectedTableIndex--
	}
	v.SetCursor(0, selectedTableIndex)
	displayTablesTree(v, tree)
	return nil
}

func nextTable(g *gocui.Gui, v *gocui.View) error {
	var tree []string
	if currentTree == "tables" {
		tree = tables
	}
	if currentTree == "columns" {
		tree = columns
	}
	if selectedTableIndex < len(tree)-1 {
		selectedTableIndex++
	}
	v.SetCursor(0, selectedTableIndex)
	displayTablesTree(v, tree)
	return nil
}

func enter(g *gocui.Gui, v *gocui.View) error {
	if currentTree == "tables" {
		currentTree = "columns"
		lastIndex = selectedTableIndex
		v.SetCursor(0, selectedTableIndex)
		return tableColumnTree(g, v)
	}
	if currentTree == "columns" {
		//TODO SELECT * FROM <Table>
		return nil
	}
	return nil
}

func escape(g *gocui.Gui, v *gocui.View) error {
	if currentTree == "columns" {
		currentTree = "tables"
		selectedTableIndex = lastIndex
		v.SetCursor(0, selectedTableIndex)
		return refreshTablesTree(g, v)
	}
	return nil
}
