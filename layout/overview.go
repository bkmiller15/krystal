package layout

import (
	"errors"
	"../printout"
	"github.com/jroimartin/gocui"
)

func overview_Layout(g *gocui.Gui) error {
	var maxX, maxY int
	var err error
	var v *gocui.View

	maxX, maxY = g.Size()

	if v, err = g.SetView("overview", -1, 1, maxX-25, maxY-25); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(tocLayout) error setting view: " + err.Error())
		}

		v.Highlight = true
		v.Wrap = false
		v.SelBgColor = gocui.ColorBlue


		if err = printout.Overview(g); err != nil {
			return err
		}


		if _, err = g.SetCurrentView("overview"); err != nil {
			return errors.New("(tocLayout) error setting current view to toc: " + err.Error())
		}
	}

	return nil
}




