package keybindings

import (
	"errors"
	"../printout"

	"github.com/jroimartin/gocui"
)

func noteTitleDeleteChar(g *gocui.Gui, v *gocui.View) error {
	if v.Editable {
		return nil
	}
	v.EditDelete(false)

	return nil
}

func noteTitleCursorLeft(g *gocui.Gui, v *gocui.View) error {
	var cx, cy, ox, oy int
	var err error

	if v.Editable {
		return nil
	}

	cx, cy = v.Cursor()

	if cx == 0 {
		return nil
	}

	if err = v.SetCursor(cx-1, cy); err != nil {
		ox, oy = v.Origin()
		if err = v.SetOrigin(ox-1, oy); err != nil {
			return errors.New("(noteTitleCursorLeft) error getting line: " + err.Error())
		}
	}
	return nil
}

func noteTitleCursorRight(g *gocui.Gui, v *gocui.View) error {
	var cx, cy, ox, oy int
	var err error
	var line string

	if v.Editable {
		return nil
	}

	cx, cy = v.Cursor()

	line, err = v.Line(cy)
	if err != nil {
		return err
	}

	if len(line) < cx+2 {
		return nil
	}

	if err = v.SetCursor(cx+1, cy); err != nil {
		ox, oy = v.Origin()
		if err = v.SetOrigin(ox+1, oy); err != nil {
			return errors.New("(noteTitleCursorRight) error getting line: " + err.Error())
		}
	}
	return nil
}

func noteTitleEnableEditable(g *gocui.Gui, v *gocui.View) error {
	if v.Editable {
		return nil
	}

	v.Editable = true

	printout.Status(g, "-- INSERT --")

	return nil
}

func noteTitleEnableEditableNextChar(g *gocui.Gui, v *gocui.View) error {
	var cx, cy, ox, oy int
	var err error

	if v.Editable {
		return nil
	}

	v.Editable = true

	printout.Status(g, "-- INSERT --")

	cx, cy = v.Cursor()

	if err = v.SetCursor(cx+1, cy); err != nil {
		ox, oy = v.Origin()
		if err = v.SetOrigin(ox, oy); err != nil {
			return errors.New("(noteTitleEnableEditableNextChar) error getting line: " + err.Error())
		}
	}

	return nil
}

func noteTitleDisableEditable(g *gocui.Gui, v *gocui.View) error {
	v.Editable = false

	printout.Status(g, "")

	return nil
}
