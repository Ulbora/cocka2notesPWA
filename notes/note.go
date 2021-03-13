package notes

import (
	"fmt"
	//lg "github.com/Ulbora/Level_Logger"
	"strconv"
	"syscall/js"

	api "github.com/Ulbora/cocka2notesApi"
)

/*

   Copyright (C) 2020 Ulbora Labs LLC. (www.ulboralabs.com)
   All rights reserved.

   Copyright (C) 2020 Ken Williamson
   All rights reserved.

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

//GetNote GetNote
func (n *NoteHandler) GetNote(this js.Value, args []js.Value) interface{} {
	//fmt.Println("this", this)
	//var id = js.ValueOf(args[0])
	var idInt = args[0].Int()
	var id = int64(idInt)
	fmt.Println("type", args[1].String())
	fmt.Println("id", id)
	go func() {
		if args[1].String() == "note" {

			document := js.Global().Get("document")
			document.Call("getElementById", "noteList").Get("style").Call("setProperty", "display", "none")
			document.Call("getElementById", "textNote").Get("style").Call("setProperty", "display", "block")
			//document.Call("getElementById", "textTitle").Set("value", note.Title)
			document.Call("getElementById", "noteId").Set("value", idInt)

			n.SaveToLocalStorage("noteType", []byte("note"))
			n.SaveToLocalStorage("noteID", []byte(strconv.FormatInt(id, 10)))

			n.PopulateTextNote(id)
			n.PollingNote = 0

			

		} else if args[1].String() == "checkbox" {

			document := js.Global().Get("document")
			document.Call("getElementById", "noteList").Get("style").Call("setProperty", "display", "none")
			document.Call("getElementById", "checkboxNote").Get("style").Call("setProperty", "display", "block")
			//document.Call("getElementById", "checkboxTitle").Set("value", note.Title)
			document.Call("getElementById", "noteId").Set("value", idInt)

			n.DoPolling = true
			n.SaveToLocalStorage("noteType", []byte("checkbox"))
			n.SaveToLocalStorage("noteID", []byte(strconv.FormatInt(id, 10)))

			n.PopulateCheckboxNote(id)

			

		}

	}()

	//noteList := n.API.GetNote(email)
	return js.Null()
}

//AddNote AddNote
func (n *NoteHandler) AddNote(this js.Value, args []js.Value) interface{} {
	document := js.Global().Get("document")
	document.Call("getElementById", "noteList").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "textNote").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "checkboxNote").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "noteUserForm").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "noteUserTable").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "loginScreen").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "changePwScreen").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "resetPwScreen").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "registerScreen").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "addNoteForm").Get("style").Call("setProperty", "display", "block")
	fmt.Println("Add a note")
	n.DoPolling = false
	n.PollingNote = 0
	return js.Null()
}

//AddNewNote AddNewNote
func (n *NoteHandler) AddNewNote(this js.Value, args []js.Value) interface{} {
	document := js.Global().Get("document")
	//document.Call("getElementById", "addNoteForm").Get("style").Call("setProperty", "display", "none")
	fmt.Println("Add a new note")
	ntitle := document.Call("getElementById", "newNoteTitle").Get("value").String()
	ntype := document.Call("getElementById", "newNoteType").Get("value").String()
	//ntype2 := document.Call("getElementById", "newNoteType2").Get("value").String()
	fmt.Println("ntitle", ntitle)
	fmt.Println("ntype", ntype)

	if ntitle != "" {
		go func() {
			var newnote api.Note
			newnote.Title = ntitle
			if ntype == "Checkbox" {
				newnote.Type = "checkbox"
			} else {
				newnote.Type = "note"
			}
			emailFn := js.Global().Get("getUserEmail")
			cemail := emailFn.Invoke()
			newnote.OwnerEmail = cemail.String()

			res := n.API.AddNote(&newnote)
			if res.Success {
				n.PopulateNoteList(cemail.String())
			}
		}()
	}

	return js.Null()
}

//DeleteNote DeleteNote
func (n *NoteHandler) DeleteNote(this js.Value, args []js.Value) interface{} {
	document := js.Global().Get("document")
	noteIDStr := document.Call("getElementById", "noteId").Get("value").String()
	fmt.Println("text note noteIDStr", noteIDStr)
	noteID, _ := strconv.ParseInt(noteIDStr, 10, 64)
	fmt.Println("text note noteID", noteID)
	emailFn := js.Global().Get("getUserEmail")
	cemail := emailFn.Invoke()
	n.DoPolling = false
	n.PollingNote = 0
	go func() {
		res := n.API.DeleteNote(noteID, cemail.String())
		if !res.Success {
			fmt.Println("Failed to delete note: ", noteID)
		}
		n.PopulateNoteList(cemail.String())
	}()

	return js.Null()
}
