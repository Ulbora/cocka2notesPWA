package notes

import (
	"fmt"
	//lg "github.com/Ulbora/Level_Logger"
	api "github.com/Ulbora/cocka2notesApi"
	//"strconv"
	"syscall/js"
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

			n.populateTextNote(id)

		} else if args[1].String() == "checkbox" {

			document := js.Global().Get("document")
			document.Call("getElementById", "noteList").Get("style").Call("setProperty", "display", "none")
			document.Call("getElementById", "checkboxNote").Get("style").Call("setProperty", "display", "block")
			//document.Call("getElementById", "checkboxTitle").Set("value", note.Title)
			document.Call("getElementById", "noteId").Set("value", idInt)

			n.pupulateCheckboxNote(id)

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
	document.Call("getElementById", "addNoteForm").Get("style").Call("setProperty", "display", "block")
	fmt.Println("Add a note")
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
			newnote.OwnerEmail = n.Email

			res := n.API.AddNote(&newnote)
			if res.Success {
				n.PopulateNoteList(n.Email)
			}
		}()
	}

	return js.Null()
}
