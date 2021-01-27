package notes

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

import (
	"fmt"
	api "github.com/Ulbora/cocka2notesApi"
	"regexp"
	"strconv"
	"syscall/js"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

//AddUsersToNote AddUsersToNote
func (n *NoteHandler) AddUsersToNote(this js.Value, args []js.Value) interface{} {
	document := js.Global().Get("document")
	noteIDStr := document.Call("getElementById", "noteId").Get("value").String()
	fmt.Println("noteIDStr", noteIDStr)
	noteID, _ := strconv.ParseInt(noteIDStr, 10, 64)
	fmt.Println(" noteID", noteID)
	document.Call("getElementById", "userworning").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "newemail").Set("value", "")
	go func() {
		note := n.API.GetCheckboxNote(noteID)
		users := n.API.GetNoteUserList(noteID, note.OwnerEmail)
		fmt.Println("users", *users)
		n.populateNoteUserScreen(noteID, users)

	}()

	return js.Null()
}

//AddNoteUser AddNoteUser
func (n *NoteHandler) AddNoteUser(this js.Value, args []js.Value) interface{} {
	document := js.Global().Get("document")
	noteIDStr := document.Call("getElementById", "noteId").Get("value").String()
	newUser := document.Call("getElementById", "newemail").Get("value").String()
	fmt.Println("noteID", noteIDStr)
	fmt.Println("newUser", newUser)
	noteID, _ := strconv.ParseInt(noteIDStr, 10, 64)
	fmt.Println(" noteID", noteID)
	if n.isEmailValid(newUser) {
		document.Call("getElementById", "notEmailWarn").Get("style").Call("setProperty", "display", "none")
		document.Call("getElementById", "userworning").Get("style").Call("setProperty", "display", "none")
		go func() {
			note := n.API.GetCheckboxNote(noteID)
			users := n.API.GetNoteUserList(noteID, note.OwnerEmail)
			var existing bool
			for _, u := range *users {
				if u == newUser {
					existing = true
				}
			}
			if !existing {
				var newusr api.NoteUsers
				newusr.NoteID = noteID
				newusr.OwnerEmail = note.OwnerEmail
				newusr.UserEmail = newUser
				res := n.API.AddUserToNote(&newusr)
				if res.Success {
					users := n.API.GetNoteUserList(noteID, note.OwnerEmail)
					n.populateNoteUserScreen(noteID, users)
				} else {
					document.Call("getElementById", "userworning").Get("style").Call("setProperty", "display", "block")
				}
			}
			n.populateNoteUserScreen(noteID, users)

		}()
	} else {
		document.Call("getElementById", "notEmailWarn").Get("style").Call("setProperty", "display", "block")
	}

	return js.Null()
}

func (n *NoteHandler) populateNoteUserScreen(noteID int64, users *[]string) {
	document := js.Global().Get("document")
	document.Call("getElementById", "textNote").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "checkboxNote").Get("style").Call("setProperty", "display", "none")

	document.Call("getElementById", "noteUserForm").Get("style").Call("setProperty", "display", "block")

	if len(*users) > 0 {
		document.Call("getElementById", "noteUserTable").Get("style").Call("setProperty", "display", "block")
		var rowHTML = ""
		for _, email := range *users {
			rowHTML = rowHTML + "<tr class='clickable-row'>"
			rowHTML = rowHTML + "<td>" + email + "</td>"
			rowHTML = rowHTML + "</tr>"
		}
		fmt.Println("rowHTML: ", rowHTML)
		document.Call("getElementById", "noteUsersBody").Set("innerHTML", rowHTML)
	}

}

func (n *NoteHandler) isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}
