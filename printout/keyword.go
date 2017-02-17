package printout

import (
	"fmt"
	//"errors"
	"../ds"
	"github.com/jroimartin/gocui"
)

func Keywords(g *gocui.Gui) error {
	var err error
	var v *gocui.View
	var index int

	if v, err = g.View("sortPromptListBox"); err != nil {
		return err
	}

	v.Clear()


	// load keyword
	ds.FilterKeywordList = ds.KeywordList


	if ds.FilterKeywordList == nil {
		fmt.Fprintf(v, " No Entries :( ")
	} else {

		for index = range ds.KeywordList {
				fmt.Fprintln(v, ds.KeywordList[index].Name)
		}
	}

	return nil
}

