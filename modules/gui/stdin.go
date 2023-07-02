package gui

import (
	"github.com/jroimartin/gocui"
)

func initKeyBinding(g *gocui.Gui) error {
	if err := kb_Tables(g); err != nil {
		return err
	}
	return nil
}

func kb_Tables(g *gocui.Gui) error {
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
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
