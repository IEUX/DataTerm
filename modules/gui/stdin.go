package gui

import (
	"github.com/jroimartin/gocui"
)

var currentView int = 0

func initKeyBinding(g *gocui.Gui) error {
	if err := kb_global(g); err != nil {
		return err
	}
	if err := kb_Navigation(g); err != nil {
		return err
	}
	if err := kb_Console(g); err != nil {
		return err
	}
	return nil
}

func kb_global(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, switchView); err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func switchView(g *gocui.Gui, v *gocui.View) error {
	var availableViews []string = []string{"Nav", "Console", "Output"}
	currentView++
	if currentView == len(availableViews) {
		currentView = 0
	}

	if _, err := g.SetCurrentView(availableViews[currentView]); err != nil {
		return err
	} else {
		v.Highlight = false
		v.Highlight = true
		g.SelFgColor = gocui.ColorMagenta
		if currentView == 1 {
			g.Cursor = true
		} else {
			g.Cursor = false
		}
	}
	return nil
}
