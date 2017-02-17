package keybindings

import (
	"errors"
	"regexp"
	"../printout"
	"../layout"
	"../ds"
	"../file"
	"strings"
	"github.com/jroimartin/gocui"
)





func prompt_Entry_Attachments(g *gocui.Gui, v *gocui.View) error {
	var err error


	if err = layout.Create_Prompt_View(g, "entryAttachments", "[Entry Attachments ]", true); err != nil {
		return err
	}

	printout.Status(g, "Update entry Attachments and press [Enter] to save.  Press [Esc] to abort.")

	return nil
}


func confirm_Prompt_Attachments(g *gocui.Gui, v *gocui.View) error {
	var attachments string
	var fileName string
	//var index int

	regexFile := regexp.MustCompile(`[/]\w+.+\w+\.\w+`)

	attachments = strings.TrimSuffix(v.ViewBuffer(), "\n")
	attachments = attachments[:len(attachments)-1]
	attachments = regexFile.FindString(attachments)
		
	index := convert_Active_Entry_To_AllEntries_Index(g)

	//ds.Set_AOF_Entry(index, aof)

	indexOverview, _ := get_Overview_Index(g)

	year, month, _, _ := ds.Get_FilteredEntry_Info(indexOverview)

	fileName = file.Copy_Attachments(attachments, year, month)

	ds.Add_Attachment_Entry(index, fileName)

	printout.Overview(g)
	printout.Content(g)

	//Write out File!
	write_Active_Entry(g)	

	abort_Prompt_Attachments(g, v)

	return nil
}



func abort_Prompt_Attachments(g *gocui.Gui, v *gocui.View) error {
	var err error

	printout.Status(g, "")


	if err = g.DeleteView("entryAttachments"); err != nil {
		return errors.New("(abortEditNoteTitle) error deleting editNoteTitle view " + err.Error())
	}

	if _, err = g.SetCurrentView("overview"); err != nil {
		return errors.New("(abortEditNoteTitle) error setting current view to toc " + err.Error())
	}

	return nil
}



func openAttachment(g *gocui.Gui, v *gocui.View) error {

	indexOverview, _ := get_Overview_Index(g)

	year, month, _, _ := ds.Get_FilteredEntry_Info(indexOverview)

	file.OpenAttachment("IMG_20160908_125329.jpg", year, month)


	return nil
}


