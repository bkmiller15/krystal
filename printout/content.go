package printout

import (
	"fmt"
	"errors"
	"../ds"
	//"../timers"
	"github.com/jroimartin/gocui"
)


func Content(g *gocui.Gui) error {
	var err error
	var index int


	Attachments(g)

	if index, err = get_Overview_Index(g); err != nil {
		return err
	}

	var contentDetailL, contentDetailR, content *gocui.View

	if err = clear_Content_Views(g); err != nil {
		return err
	}

	if ds.FilteredEntries != nil {

		if contentDetailL, contentDetailR, content, err = get_Content_Views(g); err != nil {
			return err
		}

		fmt.Fprintf(contentDetailL, " \x1b[34;4m$\x1b[0;34m Title: \x1b[0;37m%s\n", ds.FilteredEntries[index].Title)
		fmt.Fprintf(contentDetailL, " \x1b[34;4m^\x1b[0;34m Area of Focus: \x1b[0;37m%s\n", ds.FilteredEntries[index].AOF)
		fmt.Fprintf(contentDetailL, " \x1b[34;4m*\x1b[0;34m Category: \x1b[0;37m%s\n", ds.FilteredEntries[index].Category)
		fmt.Fprintf(contentDetailL, " \x1b[34;4m#\x1b[0;34m Topic: \x1b[0;37m%s\n", ds.FilteredEntries[index].Topic)
		fmt.Fprintf(contentDetailL, " \x1b[34;4m@\x1b[0;34m People: \x1b[0;37m")

		for i := range ds.FilteredEntries[index].People {
			fmt.Fprintf(contentDetailL, "%s ", ds.FilteredEntries[index].People[i] )
		}
	


		fmt.Fprintf(contentDetailR, " \x1b[34;4mD\x1b[0;34mate: \x1b[0;37m%s\n", ds.FilteredEntries[index].Year+"年 "+ds.FilteredEntries[index].Month+"月 "+ds.FilteredEntries[index].Day+"日")
		fmt.Fprintf(contentDetailR, " \x1b[34;4mE\x1b[0;34mstimated Duration Mins: \x1b[0;37m%d\n", ds.FilteredEntries[index].EstDuration)
		fmt.Fprintf(contentDetailR, " \x1b[34;4mA\x1b[0;34mctual Duration Mins: \x1b[0;37m%d\n", ds.FilteredEntries[index].Duration)


		fmt.Fprintf(content, "%s", ds.FilteredEntries[index].Content)

		if err = content.SetCursor(0, 0); err != nil {
			return errors.New("(showNoteInNoteView) error setting cursor for noteView: " + err.Error())
		}
	}

	return nil
}



func clear_Content_Views(g *gocui.Gui) error {
	var err error
	var contentDetailL, contentDetailR, content *gocui.View

	if contentDetailL, contentDetailR, content, err = get_Content_Views(g); err != nil {
		return err
	}

	contentDetailL.Clear()
	contentDetailR.Clear()
	content.Clear()

	return nil
}



func get_Content_Views(g *gocui.Gui) (*gocui.View, *gocui.View, *gocui.View, error) {
	var err error
	var contentDetailL, contentDetailR, content *gocui.View

	if contentDetailL, err = g.View("contentDetailL"); err != nil {
		return nil, nil, nil, errors.New("(getNoteViews) error grabbing handle to noteTitleView: " + err.Error())
	}

	if contentDetailR, err = g.View("contentDetailR"); err != nil {
		return nil, nil, nil, errors.New("(getNoteViews) error grabbing handle to noteNumberView: " + err.Error())
	}

	if content, err = g.View("content"); err != nil {
		return nil, nil, nil, errors.New("(getNoteViews) error grabbing handle to noteView: " + err.Error())
	}

	return contentDetailL, contentDetailR, content, nil
}







func get_Overview_Index(g *gocui.Gui) (int, error) {
	var cy, oy int 
	var err error
	var v *gocui.View

	if v, err = g.View("overview"); err != nil {
		return 0, err
	}

	_, cy = v.Cursor()
	_, oy = v.Origin()

	return (cy+oy), nil
}






