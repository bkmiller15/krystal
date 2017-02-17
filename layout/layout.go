package layout

import "github.com/jroimartin/gocui"

// Layout calls all the individual layout functions to create all the views.
func Layout(g *gocui.Gui) error {
	var err error

	if err = header_Title_Layout(g); err != nil {
		return err
	}

	if err = header_Mode_Layout(g); err != nil {
		return err
	}

	if err = overview_Layout(g); err != nil {
		return err
	}

	if err = timers_Layout(g); err != nil {
		return err
	}


	if err = content_Details_Left_Layout(g); err != nil {
		return err
	}

	if err = content_Details_Right_Layout(g); err != nil {
		return err
	}

	if err = content_Layout(g); err != nil {
		return err
	}

	if err = attachments_Layout(g); err != nil {
		return err
	}

	if err = status_Layout(g); err != nil {
		return err
	}

	return nil
}


