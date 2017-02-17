package keybindings

import (
	"../printout"
	"../ds"
	"strings"
	"github.com/jroimartin/gocui"
)


func toggle_Todos(g *gocui.Gui, v *gocui.View) error {

	if  strings.Contains(ds.ProgFilterString, "#Todos") {
		ds.ProgFilterString = strings.Replace(ds.ProgFilterString, "*Todos", "", -1)
	} else {
		ds.ProgFilterString = ds.ProgFilterString + " *Todos"
	}

	v.SetCursor(0,0)
	v.SetOrigin(0,0)

	printout.Overview(g)
	printout.Content(g)		
	
	return nil
}



