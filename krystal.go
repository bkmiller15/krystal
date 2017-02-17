package main

import (
	"log"
	"errors"
	"./layout"
	"./keybindings"
	"./file"
	"./timers"
	"./ds"
	//"time"
	//"fmt"
	
	"github.com/jroimartin/gocui"
)




func main() {
	var g *gocui.Gui
	var err error

	ds.PomTimer.Reset()

	file.Read_Config_File()

	file.Read_Log_Files()

	g = gocui.NewGui()

	if err = g.Init(); err != nil {
		log.Fatal(errors.New("(main) error initiating gui: " + err.Error()))
	}
	defer g.Close()

	g.SetLayout(layout.Layout)

	if err = keybindings.Keybindings(g); err != nil {
		log.Fatal(err)
	}



	timers.Init_Timers(g)



	g.SelBgColor = gocui.ColorBlue
	g.SelFgColor = gocui.ColorBlack
	g.Editor = &keybindings.VimEditor{}
	g.InputEsc = true
	g.Cursor = true

	if err = g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatal(err)
	}
}
