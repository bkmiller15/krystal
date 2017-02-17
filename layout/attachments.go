package layout

import (
	"fmt"
	"errors"
	"../printout"
	"github.com/jroimartin/gocui"
)


func attachments_Layout(g *gocui.Gui) error {
	var maxX, maxY int
	var v *gocui.View
	var err error

	maxX, maxY = g.Size()

	if v, err = g.SetView("attachmentsTitle", maxX-25, maxY-25, maxX, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(layout/content) error setting view: " + err.Error())
		}

		fmt.Fprintln(v, " \x1b[0;34mAttachments: \x1b[0;37m\n")

		v.Editable = false
		v.Frame = true
	}



	if v, err = g.SetView("attachments", maxX-25, maxY-24, maxX, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(layout/content) error setting view: " + err.Error())
		}

		//v.BgColor = gocui.ColorGreen
		v.Highlight = true
		v.Editable = false
		v.Frame = false

		printout.Attachments(g)


	}

	return nil
}
