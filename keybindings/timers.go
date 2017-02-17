package keybindings

import (
	"../ds"
	"../printout"
	"github.com/jroimartin/gocui"
)



func toggle_PomTimer_State(g *gocui.Gui, v *gocui.View) error {
	
	ds.PomTimer.ToggleState()
	printout.Timers(g)
	return nil
}

func reset_PomTimer(g *gocui.Gui, v *gocui.View) error {


	ds.PomTimer.Reset()
	printout.Timers(g)
	return nil
}

func reset_PomTimer_Intervals(g *gocui.Gui, v *gocui.View) error {


	ds.PomTimer.ResetIntervals()
	printout.Timers(g)
	return nil
}



func toggle_TaskTimer_State(g *gocui.Gui, v *gocui.View) error {
	
	ds.TaskTimer.ToggleState()
	printout.Timers(g)
	return nil
}

func load_TaskTimer_Selected_Entry(g *gocui.Gui, v *gocui.View) error {
	
	index := convert_Active_Entry_To_AllEntries_Index(g)

	ds.TaskTimer.LoadMin(index)
	printout.Timers(g)
	return nil
}

func save_TaskTimer_Selected_Entry(g *gocui.Gui, v *gocui.View) error {
	
	index := convert_Active_Entry_To_AllEntries_Index(g)

	ds.TaskTimer.SaveMin(index)
	printout.Timers(g)
	printout.Overview(g)
	printout.Content(g)

	//Write out File!
	write_Active_Entry(g)

	return nil
}

func reset_TaskTimer(g *gocui.Gui, v *gocui.View) error {


	ds.TaskTimer.Reset()
	printout.Timers(g)
	return nil
}

