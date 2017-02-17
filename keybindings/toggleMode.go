package keybindings

import (
	"../printout"
	"../ds"
	"strings"
	"github.com/jroimartin/gocui"
)


func toggle_Mode(g *gocui.Gui, v *gocui.View) error {


	if ds.KConfig.Mode == 0 {
		ds.KConfig.Mode = 1 
		ds.ProgFilterString = ds.ProgFilterString + " ^ALS"
	} else if ds.KConfig.Mode == 1 {
		ds.KConfig.Mode = 2
		ds.ProgFilterString = strings.Replace(ds.ProgFilterString, "^ALS", "", -1)
		ds.ProgFilterString = ds.ProgFilterString + " ^Personal"
	} else {
		ds.KConfig.Mode = 0 
		ds.ProgFilterString = strings.Replace(ds.ProgFilterString, "^Personal", "", -1)

	}

	v.SetCursor(0,0)
	v.SetOrigin(0,0)

	printout.Mode(g)
	printout.Overview(g)
	printout.Content(g)		
	
	return nil
}



