package keybindings

import (
	"errors"
	"../printout"
	"../layout"
	"../ds"
	"strings"
	"github.com/jroimartin/gocui"
)





func prompt_Entry_Title(g *gocui.Gui, v *gocui.View) error {
	var err error


	if err = layout.Create_Prompt_View(g, "entryTitle", "[Entry Title ]", true); err != nil {
		return err
	}

	printout.Status(g, "Update entry Title and press [Enter] to save.  Press [Esc] to abort.")

	return nil
}


func confirm_Prompt_Title(g *gocui.Gui, v *gocui.View) error {
	var title string
	var index int

	title = strings.TrimSuffix(v.ViewBuffer(), "\n")
	title = title[:len(title)-1]
		
	index = convert_Active_Entry_To_AllEntries_Index(g)

	ds.Set_Title_Entry(index, title)

	printout.Overview(g)
	printout.Content(g)

	//Write out File!
	write_Active_Entry(g)	

	abort_Prompt_Title(g, v)

	return nil
}



func abort_Prompt_Title(g *gocui.Gui, v *gocui.View) error {
	var err error

	printout.Status(g, "")


	if err = g.DeleteView("entryTitle"); err != nil {
		return errors.New("(abortEditNoteTitle) error deleting editNoteTitle view " + err.Error())
	}

	if _, err = g.SetCurrentView("overview"); err != nil {
		return errors.New("(abortEditNoteTitle) error setting current view to toc " + err.Error())
	}

	return nil
}





