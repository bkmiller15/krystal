package layout

import (
	"errors"
	"../printout"
	"github.com/jroimartin/gocui"
)

func timers_Layout(g *gocui.Gui) error {
	var maxX, maxY int
	var err error
	//var v *gocui.View

	maxX, maxY = g.Size()

	if _, err = g.SetView("timersview", maxX-25, 1, maxX, maxY-25); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(tocLayout) error setting view: " + err.Error())
		}


		if err = printout.Timers(g); err != nil {
			return err
		}


	}

	return nil
}




