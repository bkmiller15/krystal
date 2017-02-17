package keybindings

import (
	"errors"
	"github.com/jroimartin/gocui"
)




var trigger func(g *gocui.Gui)




func confirm_Confirmation(g *gocui.Gui, v *gocui.View) error {

	_, cy := v.Cursor()

	if cy == 1 {

		trigger(g)
	}

	



	abort_Confirmation(g, v)

	return nil
}



func abort_Confirmation(g *gocui.Gui, v *gocui.View) error {
	var err error

	//printout.Status(g, "")


	if err = g.DeleteView("confirmView"); err != nil {
		return errors.New("(abortEditNoteTitle) error deleting editNoteTitle view " + err.Error())
	}

	if _, err = g.SetCurrentView("overview"); err != nil {
		return errors.New("(abortEditNoteTitle) error setting current view to toc " + err.Error())
	}

	return nil
}





