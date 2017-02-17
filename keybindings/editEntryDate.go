package keybindings

import (
	"errors"
	"../printout"
	"../layout"
	"../ds"	
	"../file"
	"strings"
	"github.com/jroimartin/gocui"
)





func prompt_Entry_Date(g *gocui.Gui, v *gocui.View) error {
	var err error


	if err = layout.Create_Prompt_View(g, "entryDate", "[Entry Date ]", true); err != nil {
		return err
	}

	printout.Status(g, "Update entry Date and press [Enter] to save.  Press [Esc] to abort.")

	return nil
}


func confirm_Prompt_Date(g *gocui.Gui, v *gocui.View) error {
	var year, month, day string
	var date string
	var index int

	date = strings.TrimSuffix(v.ViewBuffer(), "\n")
	//date = date[:len(date)-1]	
	
	//year = date[6:len(date)-1]
	//month = date[0:2]
	//day = date[3:5]

	index = convert_Active_Entry_To_AllEntries_Index(g)

	year = ds.AllEntries[index].Year
	month = ds.AllEntries[index].Month
	day = ds.AllEntries[index].Day



	ds.Set_Date_Entry(index, date)

	printout.Overview(g)
	printout.Content(g)

	//Write Old log	
	file.Write_Log(year, month, day)

	//Write out File!
	write_Active_Entry(g)	
	
	abort_Prompt_Date(g, v)

	return nil
}



func abort_Prompt_Date(g *gocui.Gui, v *gocui.View) error {
	var err error

	printout.Status(g, "")


	if err = g.DeleteView("entryDate"); err != nil {
		return errors.New("(abortEditNoteTitle) error deleting editNoteTitle view " + err.Error())
	}

	if _, err = g.SetCurrentView("overview"); err != nil {
		return errors.New("(abortEditNoteTitle) error setting current view to toc " + err.Error())
	}

	return nil
}





