package keybindings

import (
	"errors"
	"../printout"
	"../layout"
	"../file"
	"../ds"
	"strings"
	"github.com/jroimartin/gocui"
)





func prompt_Entry_AOF(g *gocui.Gui, v *gocui.View) error {
	var err error


	if err = layout.Create_Sort_Prompt_View(g, "[Entry AOF ]"); err != nil {
		return err
	}

	printout.Status(g, "Update entry AOF and press [Enter] to save.  Press [Esc] to abort.")

	file.Read_Keyword_File("./data/aof.json") 

	printout.Keywords(g)

	return nil
}


func confirm_Prompt_AOF(g *gocui.Gui, v *gocui.View) error {
	var aof string
	var index int

	aof = strings.TrimSuffix(v.ViewBuffer(), "\n")
	aof = aof[:len(aof)-1]
		
	index = convert_Active_Entry_To_AllEntries_Index(g)

	ds.Set_AOF_Entry(index, aof)

	printout.Overview(g)
	printout.Content(g)

	//Write out File!
	write_Active_Entry(g)	

	abort_Prompt_AOF(g, v)

	return nil
}



func abort_Prompt_AOF(g *gocui.Gui, v *gocui.View) error {
	var err error

	printout.Status(g, "")

	layout.Destroy_Sort_Prompt_View(g)


	if _, err = g.SetCurrentView("overview"); err != nil {
		return errors.New("(abortEditNoteTitle) error setting current view to toc " + err.Error())
	}

	return nil
}





