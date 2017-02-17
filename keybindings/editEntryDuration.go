package keybindings

import (
	"errors"
	"../printout"
	"../layout"
	"../ds"	
	"strings"
	"strconv"
	"github.com/jroimartin/gocui"
)





func prompt_Entry_Duration(g *gocui.Gui, v *gocui.View) error {
	var err error


	if err = layout.Create_Prompt_View(g, "entryDuration", "[Entry Duration ]", true); err != nil {
		return err
	}

	printout.Status(g, "Update entry Date and press [Enter] to save.  Press [Esc] to abort.")

	return nil
}


func confirm_Prompt_Duration(g *gocui.Gui, v *gocui.View) error {
	var date string
	var dateN int
	var index int

	date = strings.TrimSuffix(v.ViewBuffer(), "\n")
	date = date[:len(date)-1]
	dateN, _ = strconv.Atoi(date)

	index = convert_Active_Entry_To_AllEntries_Index(g)

	ds.Set_Duration_Entry(index, dateN)

	printout.Overview(g)
	printout.Content(g)


	//Write out File!
	write_Active_Entry(g)	
	
	abort_Prompt_Duration(g, v)

	return nil
}



func abort_Prompt_Duration(g *gocui.Gui, v *gocui.View) error {
	var err error

	printout.Status(g, "")


	if err = g.DeleteView("entryDuration"); err != nil {
		return errors.New("(abortEditNoteTitle) error deleting editNoteTitle view " + err.Error())
	}

	if _, err = g.SetCurrentView("overview"); err != nil {
		return errors.New("(abortEditNoteTitle) error setting current view to toc " + err.Error())
	}

	return nil
}





