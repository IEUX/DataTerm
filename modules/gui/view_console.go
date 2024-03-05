package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

var (
	vbuf, buf string
)

func consoleView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	//Tables view
	if v, err := g.SetView("Console", (maxX/3)+1, 3, maxX-1, (maxY-4)/2); err != nil {

		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "SQL Console"
		v.Editable = true
		v.Wrap = true
		fmt.Fprint(v, "Console")
	}
	return nil
}
