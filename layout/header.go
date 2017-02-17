package layout

import (
	"errors"
	"fmt"
	"../printout"
	"github.com/jroimartin/gocui"
)


func header_Title_Layout(g *gocui.Gui) error {
	var err error
	var maxX int
	var v *gocui.View

	maxX, _ = g.Size()

	if v, err = g.SetView("headerTitle", maxX/2-4, -1, maxX/2+10, 1); err != nil {

		if err != gocui.ErrUnknownView {
			return errors.New("(headlineLayout) error setting view: " + err.Error())
		}

		v.Frame = false
		fmt.Fprintln(v, "\x1b[0;34mKrystal \x1b[0;37m\n")
	}

	return nil
}



func header_Mode_Layout(g *gocui.Gui) error {
	var err error
	var maxX int
	var v *gocui.View

	maxX, _ = g.Size()

	if v, err = g.SetView("headerMode", maxX-17, -1, maxX, 1); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(headlineLayout) error setting view: " + err.Error())
		}

		v.Frame = false

		printout.Mode(g)
	}

	return nil
}
