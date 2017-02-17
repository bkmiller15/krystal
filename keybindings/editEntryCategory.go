package keybindings

import (
	"errors"
	"../printout"
	"../layout"
	"../ds"
	"strings"
	"github.com/jroimartin/gocui"
)





func prompt_Entry_Category(g *gocui.Gui, v *gocui.View) error {
	var err error


	if err = layout.Create_Prompt_View(g, "entryCategory", "[Entry Category ]", true); err != nil {
		return err
	}

	printout.Status(g, "Update entry Category and press [Enter] to save.  Press [Esc] to abort.")

	return nil
}


func confirm_Prompt_Category(g *gocui.Gui, v *gocui.View) error {
	var category string
	var index int

	category = strings.TrimSuffix(v.ViewBuffer(), "\n")
	category = category[:len(category)-1]
		
	index = convert_Active_Entry_To_AllEntries_Index(g)

	ds.Set_Category_Entry(index, category)

	printout.Overview(g)
	printout.Content(g)

	//Write out File!
	write_Active_Entry(g)	
	
	abort_Prompt_Category(g, v)

	return nil
}



func abort_Prompt_Category(g *gocui.Gui, v *gocui.View) error {
	var err error

	printout.Status(g, "")


	if err = g.DeleteView("entryCategory"); err != nil {
		return errors.New("(abortEditNoteTitle) error deleting editNoteTitle view " + err.Error())
	}

	if _, err = g.SetCurrentView("overview"); err != nil {
		return errors.New("(abortEditNoteTitle) error setting current view to toc " + err.Error())
	}

	return nil
}





