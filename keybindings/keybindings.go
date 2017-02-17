package keybindings


import (
	"errors"
	"github.com/jroimartin/gocui"
)

// KeybindingsNonVim contains all the non-vim bindings.
func Keybindings(g *gocui.Gui) error {
	var err error



	// Quit application.
	if err = g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, quit); err != nil {
		return errors.New("(Keybindings) error setting keybinding for quit: " + err.Error())
	}

	// Show Hidden
	if err = g.SetKeybinding("overview", gocui.KeyCtrlH, gocui.ModNone, toggle_Show_Disabled_Entries); err != nil {
		return errors.New("(Keybindings) error setting keybinding for quit: " + err.Error())
	}

	//// Toggle Mode
	if err = g.SetKeybinding("overview", gocui.KeyTab, gocui.ModNone, toggle_Mode); err != nil {
		return errors.New("(Keybindings) error setting keybinding for quit: " + err.Error())
	}


	// Show Todos Only
	if err = g.SetKeybinding("overview", gocui.KeyF12, gocui.ModNone, toggle_Todos); err != nil {
		return errors.New("(Keybindings) error setting keybinding for quit: " + err.Error())
	}





	// Timers
	// Pomodoro Timer
	if err = g.SetKeybinding("", gocui.KeyF2, gocui.ModNone, toggle_PomTimer_State); err != nil {
		return errors.New("(Keybindings) error setting keybinding for quit: " + err.Error())
	}
	if err = g.SetKeybinding("", gocui.KeyF3, gocui.ModNone, reset_PomTimer); err != nil {
		return errors.New("(Keybindings) error setting keybinding for quit: " + err.Error())
	}
	if err = g.SetKeybinding("", gocui.KeyF4, gocui.ModNone, reset_PomTimer_Intervals); err != nil {
		return errors.New("(Keybindings) error setting keybinding for quit: " + err.Error())
	}

	// Task Timer
	if err = g.SetKeybinding("", gocui.KeyF5, gocui.ModNone, toggle_TaskTimer_State); err != nil {
		return errors.New("(Keybindings) error setting keybinding for quit: " + err.Error())
	}
	if err = g.SetKeybinding("", gocui.KeyF6, gocui.ModNone, save_TaskTimer_Selected_Entry); err != nil {
		return errors.New("(Keybindings) error setting keybinding for quit: " + err.Error())
	}
	if err = g.SetKeybinding("", gocui.KeyF7, gocui.ModNone, load_TaskTimer_Selected_Entry); err != nil {
		return errors.New("(Keybindings) error setting keybinding for quit: " + err.Error())
	}
	if err = g.SetKeybinding("", gocui.KeyF8, gocui.ModNone, reset_TaskTimer); err != nil {
		return errors.New("(Keybindings) error setting keybinding for quit: " + err.Error())
	}





	// Toggle Entry Enabled
	if err = g.SetKeybinding("overview", gocui.KeyCtrlE, gocui.ModNone, toggle_Entry_Enabled); err != nil {
		return errors.New("(Keybindings) error setting keybinding for quit: " + err.Error())
	}


	//// Create Entry.
	if err = g.SetKeybinding("", gocui.KeyCtrlN, gocui.ModNone, prompt_New_Entry_Title); err != nil {
		return errors.New("(Keybindings) error setting keybinding for quit: " + err.Error())
	}
	if err = g.SetKeybinding("newEntryTitle", gocui.KeyEnter, gocui.ModNone, prompt_New_Entry_Create); err != nil {
		return errors.New("(Keybindings) error setting keybinding for createNote: " + err.Error())
	}
	if err = g.SetKeybinding("newEntryTitle", gocui.KeyEsc, gocui.ModNone, abort_Prompt_New_Entry_Title); err != nil {
		return errors.New("(Keybindings) error setting keybinding for createNote: " + err.Error())
	}

	//Delete Entry
	if err = g.SetKeybinding("overview", 'd', gocui.ModNone, delete_Entry); err != nil {
		return errors.New("(Keybindings) error setting keybinding for showEditNoteTitle: " + err.Error())
	}








	//// Edit Entry Title
	if err = g.SetKeybinding("entryTitle", gocui.KeyEnter, gocui.ModNone, confirm_Prompt_Title); err != nil {
		return errors.New("(Keybindings) error setting keybinding for saveNoteTitle: " + err.Error())
	}
	if err = g.SetKeybinding("overview", gocui.KeyCtrl4, gocui.ModNone, prompt_Entry_Title); err != nil {
		return errors.New("(Keybindings) error setting keybinding for showEditNoteTitle: " + err.Error())
	}
	if err = g.SetKeybinding("entryTitle", gocui.KeyEsc, gocui.ModNone, abort_Prompt_Title); err != nil {
		return errors.New("(Keybindings) error setting keybinding for abortEditNoteTitle: " + err.Error())
	}


	//// Edit Entry AOF
	if err = g.SetKeybinding("sortPromptEditBox", gocui.KeyEnter, gocui.ModNone, confirm_Prompt_AOF); err != nil {
		return errors.New("(Keybindings) error setting keybinding for saveNoteTitle: " + err.Error())
	}
	if err = g.SetKeybinding("overview", gocui.KeyCtrl6, gocui.ModNone, prompt_Entry_AOF); err != nil {
		return errors.New("(Keybindings) error setting keybinding for showEditNoteTitle: " + err.Error())
	}
	if err = g.SetKeybinding("sortPromptEditBox", gocui.KeyEsc, gocui.ModNone, abort_Prompt_AOF); err != nil {
		return errors.New("(Keybindings) error setting keybinding for abortEditNoteTitle: " + err.Error())
	}

	//// Edit Entry Category
	if err = g.SetKeybinding("entryCategory", gocui.KeyEnter, gocui.ModNone, confirm_Prompt_Category); err != nil {
		return errors.New("(Keybindings) error setting keybinding for saveNoteTitle: " + err.Error())
	}
	if err = g.SetKeybinding("overview", gocui.KeyCtrl8, gocui.ModNone, prompt_Entry_Category); err != nil {
		return errors.New("(Keybindings) error setting keybinding for showEditNoteTitle: " + err.Error())
	}
	if err = g.SetKeybinding("entryCategory", gocui.KeyEsc, gocui.ModNone, abort_Prompt_Category); err != nil {
		return errors.New("(Keybindings) error setting keybinding for abortEditNoteTitle: " + err.Error())
	}


	//// Edit Entry Date
	if err = g.SetKeybinding("entryDate", gocui.KeyEnter, gocui.ModNone, confirm_Prompt_Date); err != nil {
		return errors.New("(Keybindings) error setting keybinding for saveNoteTitle: " + err.Error())
	}
	if err = g.SetKeybinding("overview", gocui.KeyCtrlD, gocui.ModNone, prompt_Entry_Date); err != nil {
		return errors.New("(Keybindings) error setting keybinding for showEditNoteTitle: " + err.Error())
	}
	if err = g.SetKeybinding("entryDate", gocui.KeyEsc, gocui.ModNone, abort_Prompt_Date); err != nil {
		return errors.New("(Keybindings) error setting keybinding for abortEditNoteTitle: " + err.Error())
	}



	//// Edit Entry Topic
	if err = g.SetKeybinding("entryTopic", gocui.KeyEnter, gocui.ModNone, confirm_Prompt_Topic); err != nil {
		return errors.New("(Keybindings) error setting keybinding for saveNoteTitle: " + err.Error())
	}
	if err = g.SetKeybinding("overview", gocui.KeyCtrl3, gocui.ModNone, prompt_Entry_Topic); err != nil {
		return errors.New("(Keybindings) error setting keybinding for showEditNoteTitle: " + err.Error())
	}
	if err = g.SetKeybinding("entryTopic", gocui.KeyEsc, gocui.ModNone, abort_Prompt_Topic); err != nil {
		return errors.New("(Keybindings) error setting keybinding for abortEditNoteTitle: " + err.Error())
	}



	//// Edit Entry Duration
	if err = g.SetKeybinding("entryDuration", gocui.KeyEnter, gocui.ModNone, confirm_Prompt_Duration); err != nil {
		return errors.New("(Keybindings) error setting keybinding for saveNoteTitle: " + err.Error())
	}
	if err = g.SetKeybinding("overview", gocui.KeyCtrlA, gocui.ModNone, prompt_Entry_Duration); err != nil {
		return errors.New("(Keybindings) error setting keybinding for showEditNoteTitle: " + err.Error())
	}
	if err = g.SetKeybinding("entryDuration", gocui.KeyEsc, gocui.ModNone, abort_Prompt_Duration); err != nil {
		return errors.New("(Keybindings) error setting keybinding for abortEditNoteTitle: " + err.Error())
	}






	//// Edit Entry Attachments
	if err = g.SetKeybinding("entryAttachments", gocui.KeyEnter, gocui.ModNone, confirm_Prompt_Attachments); err != nil {
		return errors.New("(Keybindings) error setting keybinding for saveNoteTitle: " + err.Error())
	}
	if err = g.SetKeybinding("overview", gocui.KeyCtrlP, gocui.ModNone, prompt_Entry_Attachments); err != nil {
		return errors.New("(Keybindings) error setting keybinding for showEditNoteTitle: " + err.Error())
	}
	if err = g.SetKeybinding("entryAttachments", gocui.KeyEsc, gocui.ModNone, abort_Prompt_Attachments); err != nil {
		return errors.New("(Keybindings) error setting keybinding for abortEditNoteTitle: " + err.Error())
	}


	//if err = g.SetKeybinding("", gocui.KeyCtrlL, gocui.ModNone, testBreak); err != nil {
	//	return errors.New("(Keybindings) error setting keybinding for showEditNoteTitle: " + err.Error())
	//}






	//// Search string.
	if err = g.SetKeybinding("", gocui.KeyF9, gocui.ModNone, searchString); err != nil {
		return errors.New("(Keybindings) error setting keybinding for searchString: " + err.Error())
	}
	if err = g.SetKeybinding("search", gocui.KeyEsc, gocui.ModNone, abortSearch); err != nil {
		return errors.New("(Keybindings) error setting keybinding for AbortSearch: " + err.Error())
	}
	// Find notes.
	if err = g.SetKeybinding("search", gocui.KeyEnter, gocui.ModNone, findNotes); err != nil {
		return errors.New("(Keybindings) error setting keybinding for findNotes: " + err.Error())
	}
	// Clear current filter.
	if err = g.SetKeybinding("", gocui.KeyF10, gocui.ModNone, clearCurrentFilter); err != nil {
		return errors.New("(Keybindings) error setting keybinding for clearCurrentFilter: " + err.Error())
	}



//------//------------------------------------------




	//// Navigate between toc and note views.
	if err = g.SetKeybinding("overview", gocui.KeyCtrlJ, gocui.ModNone, nextView); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for nextView: " + err.Error())
	}
	if err = g.SetKeybinding("content", gocui.KeyCtrlK, gocui.ModNone, nextView); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for nextView: " + err.Error())
	}


	if err = g.SetKeybinding("attachments", gocui.KeyCtrlH, gocui.ModNone, nextContent); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for nextView: " + err.Error())
	}
	if err = g.SetKeybinding("content", gocui.KeyCtrlL, gocui.ModNone, nextContent); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for nextView: " + err.Error())
	}


	if err = g.SetKeybinding("attachments", gocui.KeyEnter, gocui.ModNone, open_Attachment); err != nil {
		return errors.New("(Keybindings) error setting keybinding for saveNoteTitle: " + err.Error())
	}


	//// Navigate inside Overview view.
	if err = g.SetKeybinding("overview", 'j', gocui.ModNone, tocCursorDown); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for tocCursorDown: " + err.Error())
	}
	if err = g.SetKeybinding("overview", 'k', gocui.ModNone, tocCursorUp); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for tocCursorUp: " + err.Error())
	}
	if err = g.SetKeybinding("overview", gocui.KeyCtrlF, gocui.ModNone, tocPageDown); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for tocPageDown: " + err.Error())
	}
	if err = g.SetKeybinding("overview", gocui.KeyCtrlB, gocui.ModNone, tocPageUp); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for tocPageUp: " + err.Error())
	}



	// Navigate inside Attachments view.
	if err = g.SetKeybinding("attachments", 'j', gocui.ModNone, tocCursorDown); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for tocCursorDown: " + err.Error())
	}
	if err = g.SetKeybinding("attachments", 'k', gocui.ModNone, tocCursorUp); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for tocCursorUp: " + err.Error())
	}
	if err = g.SetKeybinding("attachments", gocui.KeyCtrlF, gocui.ModNone, tocPageDown); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for tocPageDown: " + err.Error())
	}
	if err = g.SetKeybinding("attachments", gocui.KeyCtrlB, gocui.ModNone, tocPageUp); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for tocPageUp: " + err.Error())
	}





	// Confirmation View
	if err = g.SetKeybinding("confirmView", 'j', gocui.ModNone, tocCursorDown); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for tocCursorDown: " + err.Error())
	}
	if err = g.SetKeybinding("confirmView", 'k', gocui.ModNone, tocCursorUp); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for tocCursorUp: " + err.Error())
	}
	if err = g.SetKeybinding("confirmView", gocui.KeyEsc, gocui.ModNone, abort_Confirmation); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for tocPageDown: " + err.Error())
	}
	if err = g.SetKeybinding("confirmView", gocui.KeyEnter, gocui.ModNone, confirm_Confirmation); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for tocPageUp: " + err.Error())
	}



	//// Save a record.
	//if err = g.SetKeybinding("noteDetail", gocui.KeyCtrlW, gocui.ModNone, saveNote); err != nil {
	//	return errors.New("(KeybindingsVim) error setting keybinding for saveNote: " + err.Error())
	//}

	//// New record.
	//if err = g.SetKeybinding("toc", gocui.KeyCtrlI, gocui.ModNone, newRec); err != nil {
	//	return errors.New("(KeybindingsVim) error setting keybinding for newRec: " + err.Error())
	//}
	//if err = g.SetKeybinding("noteDetail", gocui.KeyCtrlI, gocui.ModNone, newRec); err != nil {
	//	return errors.New("(KeybindingsVim) error setting keybinding for newRec: " + err.Error())
	//}
	//if err = g.SetKeybinding("newTitle", gocui.KeyEsc, gocui.ModNone, abortNewTitle); err != nil {
	//	return errors.New("(KeybindingsVim) error setting keybinding for abortNewTitle: " + err.Error())
	//}

	//// Navigation inside note view
	//if err = g.SetKeybinding("content", '0', gocui.ModNone, noteBeginningOfLine); err != nil {
	//	return errors.New("(KeybindingsVim) error setting keybinding for noteBeginningOfLine: " + err.Error())
	//}
	//if err = g.SetKeybinding("content", '$', gocui.ModNone, noteEndOfLine); err != nil {
	//	return errors.New("(KeybindingsVim) error setting keybinding for noteEndOfLine: " + err.Error())
	//}
	//if err = g.SetKeybinding("content", gocui.KeyCtrlF, gocui.ModNone, notePageDown); err != nil {
	//	return errors.New("(KeybindingsVim) error setting keybinding for notePageDown: " + err.Error())
	//}
	//if err = g.SetKeybinding("content", gocui.KeyCtrlB, gocui.ModNone, notePageUp); err != nil {
	//	return errors.New("(KeybindingsVim) error setting keybinding for notePageUp: " + err.Error())
	//}
	//if err = g.SetKeybinding("content", gocui.KeyEsc, gocui.ModNone, noteDisableEditable); err != nil {
	//	return errors.New("(KeybindingsVim) error setting keybinding for noteDisableEditable: " + err.Error())
	//}
	//if err = g.SetKeybinding("content", 'h', gocui.ModNone, noteCursorLeft); err != nil {
	//	return errors.New("(KeybindingsVim) error setting keybinding for noteCursorLeft: " + err.Error())
	//}
	//if err = g.SetKeybinding("content", 'i', gocui.ModNone, noteEnableEditable); err != nil {
	//	return errors.New("(KeybindingsVim) error setting keybinding for noteEnableEditable: " + err.Error())
	//}
	//if err = g.SetKeybinding("content", 'j', gocui.ModNone, noteCursorDown); err != nil {
	//	return errors.New("(KeybindingsVim) error setting keybinding for noteCursorDown: " + err.Error())
	//}
	//if err = g.SetKeybinding("content", 'k', gocui.ModNone, noteCursorUp); err != nil {
	//	return errors.New("(KeybindingsVim) error setting keybinding for noteCursorUp: " + err.Error())
	//}
	//if err = g.SetKeybinding("content", 'l', gocui.ModNone, noteCursorRight); err != nil {
	//	return errors.New("(KeybindingsVim) error setting keybinding for noteCursorRight: " + err.Error())
	//}
	//if err = g.SetKeybinding("content", 'a', gocui.ModNone, noteEnableEditableNextChar); err != nil {
	//	return errors.New("(KeybindingsVim) error setting keybinding for noteEnableEditableNextChar: " + err.Error())
	//}
	//if err = g.SetKeybinding("content", 'x', gocui.ModNone, noteDeleteChar); err != nil {
	//	return errors.New("(KeybindingsVim) error setting keybinding for noteDeleteChar: " + err.Error())
	//}





	//// Navigation inside edit note title view
	if err = g.SetKeybinding("editEntryTitle", 'a', gocui.ModNone, noteTitleEnableEditableNextChar); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteTitleEnableEditableNextChar: " + err.Error())
	}
	if err = g.SetKeybinding("editEntryTitle", 'h', gocui.ModNone, noteTitleCursorLeft); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteTitleCursorLeft: " + err.Error())
	}
	//if err = g.SetKeybinding("editEntryTitle", 'i', gocui.ModNone, noteTitleEnableEditable); err != nil {
	//	return errors.New("(KeybindingsVim) error setting keybinding for noteTitleEnableEditable: " + err.Error())
	//}
	if err = g.SetKeybinding("editEntryTitle", 'l', gocui.ModNone, noteTitleCursorRight); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteTitleCursorRight: " + err.Error())
	}
	if err = g.SetKeybinding("editEntryTitle", 'x', gocui.ModNone, noteTitleDeleteChar); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteTitleDeleteChar: " + err.Error())
	}
	//if err = g.SetKeybinding("editEntryTitle", gocui.KeyEsc, gocui.ModNone, noteTitleDisableEditable); err != nil {
	//	return errors.New("(KeybindingsVim) error setting keybinding for noteTitleDisableEditable: " + err.Error())
	//}

	return nil
}
