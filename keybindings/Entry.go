package keybindings

import (
	"errors"
	"../printout"
	"../layout"
	"../ds"
	"../helper"
	"../file"
	"strings"
	"github.com/jroimartin/gocui"
)





func prompt_New_Entry_Title(g *gocui.Gui, v *gocui.View) error {
	var err error


	if err = layout.Create_Prompt_View(g, "newEntryTitle", "[ New Entry Title ]", true); err != nil {
		return err
	}



	printout.Status(g, "Update note title and press [Enter] to save.  Press [Esc] to abort.")

	return nil
}


func prompt_New_Entry_Create(g *gocui.Gui, v *gocui.View) error {
	var err error
	var year, month, day, title string
	var overview *gocui.View


	
	title = strings.TrimSuffix(v.ViewBuffer(), "\n")
	title = title[:len(title)-1]

	year, month, day = helper.Get_Current_Date()
		

	ds.Add_New_Entry(year, month, day, title)


	if _, err = g.SetCurrentView("content"); err != nil {
		return errors.New("(saveNoteTitle) error setting current view to toc " + err.Error())
	}

	if err = g.DeleteView("newEntryTitle"); err != nil {
		return errors.New("(saveNoteTitle) error deleting editNoteTitle view " + err.Error())
	}



	if overview, err = g.SetCurrentView("overview"); err != nil {
		return errors.New("(findNotes) error setting current view to toc: " + err.Error())
	}

	ds.UserFilterString = ""

	overview.SetCursor(0,0)
	overview.SetOrigin(0,0)

	printout.Overview(g)
	printout.Content(g)

	//Write out File!
	write_Active_Entry(g)

	return nil
}



func abort_Prompt_New_Entry_Title(g *gocui.Gui, v *gocui.View) error {
	var err error

	printout.Status(g, "")


	if err = g.DeleteView("newEntryTitle"); err != nil {
		return errors.New("(abortEditNoteTitle) error deleting editNoteTitle view " + err.Error())
	}

	if _, err = g.SetCurrentView("overview"); err != nil {
		return errors.New("(abortEditNoteTitle) error setting current view to toc " + err.Error())
	}

	return nil
}





func delete_Entry(g *gocui.Gui, v *gocui.View) error {


	trigger = func(g *gocui.Gui) {

		Findex, _ := get_Overview_Index(g)

		year, month, day, number := ds.Get_FilteredEntry_Info(Findex)

		index := ds.Get_Entry_Index(year, month, day, number)

		ds.Delete_Entry(index)

		file.Write_Log(year, month, day)

		printout.Overview(g)
		printout.Content(g)
	}

	layout.Create_Confirmation_View(g, "[Delete Item ?]")

	return nil

}

