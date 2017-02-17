package layout
import (
	"errors"
	"github.com/jroimartin/gocui"
)

func Create_Break_View(g *gocui.Gui) error {
	var err error
	var maxX, maxY int
	var v *gocui.View

	maxX, maxY = g.Size()

	if v, err = g.SetView("breakScreen",  -1, 1, maxX-25, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(createInputView) error creating view: " + err.Error())
		}

		v.BgColor = gocui.ColorGreen


		if _, err = g.SetCurrentView("breakScreen"); err != nil {
			return errors.New("(createInputView) error setting current view: " + err.Error())
		}
	}
	return nil
}




