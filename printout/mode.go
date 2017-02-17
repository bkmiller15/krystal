package printout

import (
	"fmt"
	"../ds"
	"github.com/jroimartin/gocui"
)

func Mode(g *gocui.Gui) error {
	var err error
	var headerView *gocui.View

	if headerView, err = g.View("headerMode"); err != nil {
		return err
	}

	headerView.Clear()
	if ds.KConfig.Mode == 1 {
		fmt.Fprint(headerView, "\x1b[0;34m     Mode: Work\x1b[0;37m")
	} else if ds.KConfig.Mode == 2 {
		fmt.Fprint(headerView, "\x1b[0;34m Mode: Personal\x1b[0;37m")
	} else {
		fmt.Fprint(headerView, "\x1b[0;34m      Mode: All\x1b[0;37m")
	}
	


	return nil
}

