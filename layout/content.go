package layout

import (
	"errors"
	"../printout"
	"github.com/jroimartin/gocui"
)

func content_Details_Left_Layout(g *gocui.Gui) error {
	var maxX, maxY int
	var v *gocui.View
	var err error

	maxX, maxY = g.Size()

	if v, err = g.SetView("contentDetailL", -1, maxY-25, maxX-55, maxY-19); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(layout/content) error setting view: " + err.Error())
		}
		v.Frame = false
	}

	return nil
}

func content_Details_Right_Layout(g *gocui.Gui) error {
	var maxX, maxY int
	var v *gocui.View
	var err error

	maxX, maxY = g.Size()

	if v, err = g.SetView("contentDetailR", maxX-56, maxY-25, maxX-25, maxY-19); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(layout/content) error setting view: " + err.Error())
		}
		v.Frame = false
	}

	return nil
}

func content_Layout(g *gocui.Gui) error {
	var maxX, maxY int
	var v *gocui.View
	var err error

	maxX, maxY = g.Size()

	if v, err = g.SetView("content", 0, maxY-19, maxX-26, maxY-3); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(layout/content) error setting view: " + err.Error())
		}


		printout.Content(g)

		//v.Editable = false
		v.Editable = true
		v.Wrap = false
		v.Frame = false
	}

	return nil
}
