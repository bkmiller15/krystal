package printout

import (
	"fmt"
	//"errors"
	"github.com/jroimartin/gocui"
)

func Status(g *gocui.Gui, msg string) error {
	var err error
	var statusView *gocui.View

	if statusView, err = g.View("status"); err != nil {
		return err
	}

	statusView.Clear()
	fmt.Fprint(statusView, " ")
	fmt.Fprint(statusView, msg)

	return nil
}

