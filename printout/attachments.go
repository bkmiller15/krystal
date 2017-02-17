package printout

import (
	"fmt"
	//"errors"
	"strings"
	"../ds"
	"github.com/jroimartin/gocui"
)


func Attachments(g *gocui.Gui) error {
	var err error
	var attachmentsView *gocui.View
	var index int

	if index, err = get_Overview_Index(g); err != nil {
		return err
	}

	if attachmentsView, err = g.View("attachments"); err != nil {
		return err
	}

	attachmentsView.Clear()

	if ds.FilteredEntries != nil {
		for i := range ds.FilteredEntries[index].Attachments {
			fmt.Fprintf(attachmentsView, " ")
			fmt.Fprintln(attachmentsView, ds.FilteredEntries[index].Attachments[i]+ strings.Repeat(" ", 50))
		}
	}

	return nil
}






