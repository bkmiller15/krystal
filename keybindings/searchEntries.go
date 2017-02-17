package keybindings

import (
	"errors"
	"strings"
	"../ds"
	"../layout"
	"../printout"
	"github.com/jroimartin/gocui"
)

func abortSearch(g *gocui.Gui, v *gocui.View) error {
	var err error

	printout.Status(g, "")

	if err = g.DeleteView("search"); err != nil {
		return errors.New("(AbortSearch) error deleting search view " + err.Error())
	}

	if _, err = g.SetCurrentView("overview"); err != nil {
		return errors.New("(AbortSearch) error setting current view to toc " + err.Error())
	}
	return nil
}

func clearCurrentFilter(g *gocui.Gui, v *gocui.View) error {
	var err error
	var overview *gocui.View

	if overview, err = g.SetCurrentView("overview"); err != nil {
		return errors.New("(findNotes) error setting current view to toc: " + err.Error())
	}

	ds.UserFilterString = ""

	overview.SetCursor(0,0)
	overview.SetOrigin(0,0)

	printout.Overview(g)
	printout.Content(g)

	return nil
}

func findNotes(g *gocui.Gui, v *gocui.View) error {
	var err error
	var searchString string
	var searchView *gocui.View
	var overview *gocui.View

	if searchView, err = g.View("search"); err != nil {
		return errors.New("(findNotes) error grabbing handle to search view: " + err.Error())
	}

	searchString = strings.TrimSuffix(searchView.ViewBuffer(), "\n")
	searchString = searchString[:len(searchString)-1]

	ds.UserFilterString = searchString

	if err = g.DeleteView("search"); err != nil {
		return errors.New("(findNotes) error deleting search view " + err.Error())
	}

	

	if overview, err = g.SetCurrentView("overview"); err != nil {
		return errors.New("(findNotes) error setting current view to toc: " + err.Error())
	}

	overview.SetCursor(0,0)
	overview.SetOrigin(0,0)

	printout.Overview(g)
	printout.Content(g)


	return nil
}

func searchString(g *gocui.Gui, v *gocui.View) error {
	var err error

	if err = layout.Create_Prompt_View(g, "search", "Enter word(s) to search on:", true); err != nil {
		return err
	}

	printout.Status(g, "Enter a search string and press [Enter] to search.  Press [Esc] to abort.")

	return nil
}





