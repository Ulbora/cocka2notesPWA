package notes

import (
	"fmt"
	lg "github.com/Ulbora/Level_Logger"
	api "github.com/Ulbora/cocka2notesApi"
	"strconv"
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

//Handler Handler
type Handler interface {
	Login(this js.Value, args []js.Value) interface{}

	GetNoteList(this js.Value, args []js.Value) interface{}
	GetNote(this js.Value, args []js.Value) interface{}
	AddNote(this js.Value, args []js.Value) interface{}
	AddNewNote(this js.Value, args []js.Value) interface{}

	UpdateCheckboxNoteTitle(this js.Value, args []js.Value) interface{}
	UpdateCheckboxNoteItem(this js.Value, args []js.Value) interface{}
	AddCheckboxNoteItem(this js.Value, args []js.Value) interface{}
	DeleteCheckboxNoteItem(this js.Value, args []js.Value) interface{}

	UpdateTextNoteTitle(this js.Value, args []js.Value) interface{}
	UpdateTextNoteItem(this js.Value, args []js.Value) interface{}
	AddTextNoteItem(this js.Value, args []js.Value) interface{}
	DeleteTextNoteItem(this js.Value, args []js.Value) interface{}

	AddUsersToNote(this js.Value, args []js.Value) interface{}
	AddNoteUser(this js.Value, args []js.Value) interface{}
}

//NoteHandler NoteHandler
type NoteHandler struct {
	API   api.API
	Log   *lg.Logger
	Email string
}

//GetNew GetNew
func (n *NoteHandler) GetNew() Handler {
	return n
}

//GetNoteList GetNoteList
func (n *NoteHandler) GetNoteList(this js.Value, args []js.Value) interface{} {
	n.PopulateNoteList(n.Email)
	return js.Null()
}

//Login Login
func (n *NoteHandler) Login(this js.Value, args []js.Value) interface{} {
	go func() {
		document := js.Global().Get("document")
		email := document.Call("getElementById", "email").Get("value").String()
		fmt.Println(email)
		pw := document.Call("getElementById", "password").Get("value").String()
		fmt.Println(pw)
		var u api.User
		u.Email = email
		u.Password = pw
		res := n.API.Login(&u)
		fmt.Println("login suc: ", *res)
		if res.Success {
			emailc := js.Global().Get("setUserEmail")
			emailc.Invoke(email)
			n.Email = email
			document := js.Global().Get("document")
			document.Call("getElementById", "loginScreen").Get("style").Call("setProperty", "display", "none")
			//go to note list
			n.PopulateNoteList(email)
		} else if res.Email == "" {
			// auto create a new account
			suc := n.API.AddUser(&u)
			if suc.Success {
				emailc := js.Global().Get("setUserEmail")
				emailc.Invoke(email)
				document := js.Global().Get("document")
				document.Call("getElementById", "loginScreen").Get("style").Call("setProperty", "display", "none")
				// gto to note list
				n.PopulateNoteList(email)
			}
		}

	}()
	return js.Null()
}

//PopulateNoteList PopulateNoteList
func (n *NoteHandler) PopulateNoteList(email string) {
	go func() {
		noteList := n.API.GetUsersNotes(email)
		fmt.Println("Note list in PopulateNoteList", *noteList)
		document := js.Global().Get("document")
		document.Call("getElementById", "noteList").Get("style").Call("setProperty", "display", "block")
		document.Call("getElementById", "checkboxNote").Get("style").Call("setProperty", "display", "none")
		document.Call("getElementById", "textNote").Get("style").Call("setProperty", "display", "none")
		//document.Call("getElementById", "checkboxNote").Get("style").Call("setProperty", "display", "block")
		document.Call("getElementById", "noteUserForm").Get("style").Call("setProperty", "display", "none")
		document.Call("getElementById", "noteUserTable").Get("style").Call("setProperty", "display", "none")
		document.Call("getElementById", "addNoteForm").Get("style").Call("setProperty", "display", "none")
		document.Call("getElementById", "addNoteForm").Get("style").Call("setProperty", "display", "none")
		document.Call("getElementById", "newNoteTitle").Set("value", "")
		var rowHTML = ""
		for i, nt := range *noteList {
			var row = i + 1
			fmt.Println("note: ", nt)
			var rowStr = strconv.Itoa(row)
			var idStr = strconv.FormatInt(nt.ID, 10)
			var rowTime = nt.LastUsed.Format("2006-01-02 15:04:05")
			rowHTML = rowHTML + "<tr class='clickable-row' onclick='showNote(" + idStr + ",\"" + nt.Type + "\")'>"
			rowHTML = rowHTML + "<th scope='row'>" + rowStr + "</th>"
			rowHTML = rowHTML + "<td>" + nt.Title + "</td>"
			rowHTML = rowHTML + "<td>" + rowTime + "</td>"
			rowHTML = rowHTML + "</tr>"

		}
		fmt.Println("rowHTML: ", rowHTML)
		document.Call("getElementById", "noteListBody").Set("innerHTML", rowHTML)
		// document.Call("getElementById", "job").Set("value", (*noteList)[0].Title)
		//email := js.Global().Get("document").Get("email").String()
		//fmt.Println(email)
	}()
}
