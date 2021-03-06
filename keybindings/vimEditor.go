package keybindings

import (
	"errors"
	"github.com/jroimartin/gocui"
)

type VimEditor struct {
	Insert bool
	DeleteMode bool
}

func (ve *VimEditor) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	if ve.Insert {
		ve.InsertMode(v, key, ch, mod)
		return
	}
	ve.NormalMode(v, key, ch, mod)
}

func (ve *VimEditor) InsertMode(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case key == gocui.KeyEsc:
		ve.Insert = false
	case ch != 0 && mod == 0:
		v.EditWrite(ch)
	case key == gocui.KeySpace:
		v.EditWrite(' ')
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		v.EditDelete(true)
	case key == gocui.KeyDelete:
		v.EditDelete(false)
	case key == gocui.KeyInsert:
		v.Overwrite = !v.Overwrite
	case key == gocui.KeyEnter:
		v.EditNewLine()
	case key == gocui.KeyArrowDown:
		v.MoveCursor(0, 1, false)
	case key == gocui.KeyArrowUp:
		v.MoveCursor(0, -1, false)
	case key == gocui.KeyArrowLeft:
		v.MoveCursor(-1, 0, false)
	case key == gocui.KeyArrowRight:
		v.MoveCursor(1, 0, false)
	}
	// TODO: handle other keybindings...
}

func (ve *VimEditor) NormalMode(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case key == gocui.KeyEsc:
		ve.Insert = false
		ve.DeleteMode = false
     	case ch == '0':
		_, cy := v.Cursor()
		v.SetCursor(0, cy)
		_, oy := v.Origin()
		v.SetOrigin(0, oy)
        case ch == '$':
               vimEndOfLine(v)
	case ch == 'i':
		ve.Insert = true
        case ch == 'j':
                v.MoveCursor(0, 1, false)
        case ch == 'k':
                v.MoveCursor(0, -1, false)
        case ch == 'h':
                v.MoveCursor(-1, 0, false)
        case ch == 'l':
                v.MoveCursor(1, 0, false)
        case ch == 'd':
		if ve.DeleteMode {
			ve.DeleteMode = false
		} else {
			ve.DeleteMode = true
		}
        case ch == 'x':
               v.EditDelete(false)
        case ch == 'o':
		v.MoveCursor(0, 1, false)
                v.EditNewLine()
		v.MoveCursor(0, -1, false)
		ve.Insert = true

  
	}
	// TODO: handle other keybindings...
}





func vimBeginningOfLine(v *gocui.View) error {
	var cy, oy int
	var err error

	if v.Editable {
		return nil
	}

	_, cy = v.Cursor()
	v.SetCursor(0, cy)

	if err = v.SetCursor(0, cy); err != nil {
		_, oy = v.Origin()
		if err = v.SetOrigin(0, oy); err != nil {
			return errors.New("(noteBeginningOfLine) error setting origin: " + err.Error())
		}
	}
	return nil
}

func vimEndOfLine(v *gocui.View) error {
	var cy, oy int
	var err error
	var lineLength int
	var line string

	if v.Editable {
		return nil
	}

	_, cy = v.Cursor()

	if line, err = v.Line(cy); err != nil {
		return errors.New("(noteEndOfLine) error getting line: " + err.Error())
	}

	lineLength = len(line)

	if lineLength > 0 {
		lineLength -= 1
	}

	if err = v.SetCursor(lineLength, cy); err != nil {
		_, oy = v.Origin()
		if err = v.SetOrigin(lineLength, oy); err != nil {
			return errors.New("(noteEndOfLine) error setting origin: " + err.Error())
		}
	}
	return nil
}


