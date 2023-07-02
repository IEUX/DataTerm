package gui

import (
	"github.com/jroimartin/gocui"
)

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
		selectedTableIndex = 0
		v.SetCursor(0, selectedTableIndex)
		return tableColumnTreeView(g, v)
	}
	if currentTree == "columns" {
		return nil
	}
	return nil
}

func escape(g *gocui.Gui, v *gocui.View) error {
	if currentTree == "columns" {
		currentTree = "tables"
		selectedTableIndex = lastIndex
		v.SetCursor(0, selectedTableIndex)
		return refreshTablesTreeView(g, v)
	}
	return nil

}

// func openTable(g *gocui.Gui, v *gocui.View) error {
// 	g.Update(tableView)
// 	return nil
// }
