package layout
import (
	"errors"
	"github.com/jroimartin/gocui"
)

func Create_Prompt_View(g *gocui.Gui, vName string, vTitle string, editable bool) error {
	var err error
	var maxX, maxY int
	var v *gocui.View

	maxX, maxY = g.Size()

	if v, err = g.SetView(vName, maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(createInputView) error creating view: " + err.Error())
		}

		v.Title = vTitle
		v.Editable = editable

		if _, err = g.SetCurrentView(vName); err != nil {
			return errors.New("(createInputView) error setting current view: " + err.Error())
		}
	}
	return nil
}



func Create_Sort_Prompt_View(g *gocui.Gui, vTitle string) error {
	var err error
	var maxX, maxY int
	var v *gocui.View

	maxX, maxY = g.Size()

	if v, err = g.SetView("sortPromptMain", maxX/2-30, maxY/2-5, maxX/2+30, maxY/2+10); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(createInputView) error creating view: " + err.Error())
		}

		v.Title = vTitle

		if _, err = g.SetCurrentView("sortPromptMain"); err != nil {
			return errors.New("(createInputView) error setting current view: " + err.Error())
		}
	}

	if v, err = g.SetView("sortPromptEditBox", maxX/2-30, maxY/2-5, maxX/2+30, maxY/2-3); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(createInputView) error creating view: " + err.Error())
		}

		v.BgColor = gocui.ColorRed               
		v.Editable = true
		v.Wrap = false
		
	}
	if v, err = g.SetView("sortPromptListBox", maxX/2-30, maxY/2-3, maxX/2+30, maxY/2+10); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(createInputView) error creating view: " + err.Error())
		}

		v.BgColor = gocui.ColorYellow  
	}

	if _, err = g.SetCurrentView("sortPromptEditBox"); err != nil {
		return errors.New("(abortEditNoteTitle) error setting current view to toc " + err.Error())
	}


	return nil
}




func Destroy_Sort_Prompt_View(g *gocui.Gui) {
	var err error 

	if err = g.DeleteView("sortPromptMain"); err != nil {
		//return errors.New("(abortEditNoteTitle) error deleting editNoteTitle view " + err.Error())
	}
	if err = g.DeleteView("sortPromptEditBox"); err != nil {
		//return errors.New("(abortEditNoteTitle) error deleting editNoteTitle view " + err.Error())
	}
	if err = g.DeleteView("sortPromptListBox"); err != nil {
		//return errors.New("(abortEditNoteTitle) error deleting editNoteTitle view " + err.Error())
	}
}




