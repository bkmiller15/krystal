package keybindings

import (
	"../ds"
	"../printout"
	"strings"
	"github.com/jroimartin/gocui"
)


func nextView(g *gocui.Gui, v *gocui.View) error {
	var content string
	var index int
	var err error

	if v == nil || v.Name() == "overview" {
		if _, err = g.SetCurrentView("content"); err != nil {
			return err
		}
	} else {
		content = strings.TrimSuffix(v.ViewBuffer(), "\n")
		
		index = convert_Active_Entry_To_AllEntries_Index(g)
		
		ds.Set_Content_Entry(index, content)

		printout.Overview(g)
		printout.Content(g)

		//Write out File!
		write_Active_Entry(g)

		if _, err = g.SetCurrentView("overview"); err != nil {
			return err
		}
	}



	return nil
}







func nextContent(g *gocui.Gui, v *gocui.View) error {
	var err error

	if v == nil || v.Name() == "content" {
		if _, err = g.SetCurrentView("attachments"); err != nil {
			return err
		}
	} else {
		if _, err = g.SetCurrentView("content"); err != nil {
			return err
		}
	}



	return nil
}



func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
