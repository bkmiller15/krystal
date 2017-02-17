package keybindings


import (
	//"errors"
	"../ds"
	"../printout"
	"../file"
	"../layout"
	"github.com/jroimartin/gocui"
)

func get_Overview_Index(g *gocui.Gui) (int, error) {
	var cy, oy int 
	var err error
	var v *gocui.View

	if v, err = g.View("overview"); err != nil {
		return 0, err
	}

	_, cy = v.Cursor()
	_, oy = v.Origin()

	return (cy+oy), nil
}



func write_Active_Entry(g *gocui.Gui) {
	var i int

	i, _ = get_Overview_Index(g)

	file.Write_Log(ds.FilteredEntries[i].Year, ds.FilteredEntries[i].Month, ds.FilteredEntries[i].Day)
}


func convert_Active_Entry_To_AllEntries_Index(g *gocui.Gui) int {
	var indexOverview, number int
	var year, month, day string

	indexOverview, _ = get_Overview_Index(g)

	year, month, day, number = ds.Get_FilteredEntry_Info(indexOverview)

	return ds.Get_Entry_Index(year, month, day, number)

}


func toggle_Show_Disabled_Entries(g *gocui.Gui, v *gocui.View) error {

	if ds.KConfig.ShowDisabled { ds.KConfig.ShowDisabled = false } else { ds.KConfig.ShowDisabled = true }

	v.SetCursor(0,0)
	v.SetOrigin(0,0)

	printout.Overview(g)
	printout.Content(g)

	return nil
}



func testBreak(g *gocui.Gui, v *gocui.View) error {

	layout.Create_Break_View(g)
	return nil
}
















