package layout
import (
	"fmt"
	"errors"
	"strings"
	"github.com/jroimartin/gocui"
)



func Create_Confirmation_View(g *gocui.Gui, vTitle string) error {
	var err error
	var maxX, maxY int
	var v *gocui.View

	maxX, maxY = g.Size()

	if v, err = g.SetView("confirmView", maxX/2-30, maxY/2, maxX/2+30, maxY/2+3); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(createInputView) error creating view: " + err.Error())
		}

		v.Title = vTitle
		v.Highlight = true

		
		fmt.Fprintln(v, " No - Never Mind" + strings.Repeat(" ", 60))
		fmt.Fprintln(v, " Yes - Go Ahead" + strings.Repeat(" ", 60))		
		if _, err = g.SetCurrentView("confirmView"); err != nil {
			return errors.New("(createInputView) error setting current view: " + err.Error())
		}
	}
	return nil
}






