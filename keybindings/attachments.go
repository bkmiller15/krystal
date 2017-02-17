package keybindings

import (
	"../file"
	"../ds"
	"github.com/jroimartin/gocui"
)





func open_Attachment(g *gocui.Gui, v *gocui.View) error {
	var cy, oy int 

	_, cy = v.Cursor()
	_, oy = v.Origin()
	

	indexOverview, _ := get_Overview_Index(g)

	year, month, _, _ := ds.Get_FilteredEntry_Info(indexOverview)


	file.OpenAttachment(ds.FilteredEntries[indexOverview].Attachments[(cy+oy)], year, month)

	return nil
}






