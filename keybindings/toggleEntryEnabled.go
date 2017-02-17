package keybindings

import (
	"../printout"
	"../ds"
	"github.com/jroimartin/gocui"
)


func toggle_Entry_Enabled(g *gocui.Gui, v *gocui.View) error {
	var index int
		
	index = convert_Active_Entry_To_AllEntries_Index(g)

	ds.Set_Toggle_Enabled_Entry(index)

	printout.Overview(g)
	printout.Content(g)

	//Write out File!
	write_Active_Entry(g)		
	
	return nil
}



