package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func outputView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	//Tables view
	if v, err := g.SetView("Output", (maxX/3)+1, maxY-maxY/2-2, maxX-1, maxY-1); err != nil {

		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Result"
		v.Wrap = true
		fmt.Fprint(v, "Output")
	}
	return nil
}
