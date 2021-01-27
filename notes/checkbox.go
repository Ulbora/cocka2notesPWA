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
	//lg "github.com/Ulbora/Level_Logger"
	api "github.com/Ulbora/cocka2notesApi"
	"strconv"
	"syscall/js"
)

//UpdateCheckboxNoteItem UpdateCheckboxNoteItem
func (n *NoteHandler) UpdateCheckboxNoteItem(this js.Value, args []js.Value) interface{} {

	fmt.Println("id", args[0])
	fmt.Println("noteId", args[1])
	fmt.Println("checked", args[2])
	fmt.Println("text", args[3])

	var idInt = args[0].Int()
	var id = int64(idInt)
	var idStr = strconv.FormatInt(id, 10)
	var noteIDInt = args[1].Int()
	var noteID = int64(noteIDInt)
	//var txt = args[3].String()
	document := js.Global().Get("document")
	txt := document.Call("getElementById", idStr).Get("value").String()
	go func() {
		var item api.CheckboxNoteItem
		item.ID = id
		item.NoteID = noteID
		item.Text = txt
		if args[2].String() == "checked" {
			item.Checked = false
		} else {
			item.Checked = true
		}
		res := n.API.UpdateCheckboxItem(&item)
		fmt.Println("cb update suc: ", res.Success)

		n.pupulateCheckboxNote(noteID)

	}()

	return js.Null()
}

//AddCheckboxNoteItem AddCheckboxNoteItem
func (n *NoteHandler) AddCheckboxNoteItem(this js.Value, args []js.Value) interface{} {
	fmt.Println("noteId", args[0])
	var noteIDInt = args[0].Int()
	var noteID = int64(noteIDInt)
	fmt.Println("noteId", noteID)
	document := js.Global().Get("document")
	txt := document.Call("getElementById", "newtxt").Get("value").String()
	fmt.Println("txt", txt)
	go func() {
		var item api.CheckboxNoteItem
		item.NoteID = noteID
		item.Text = txt
		res := n.API.AddCheckboxItem(&item)
		fmt.Println("cb update suc: ", res.Success)

		n.pupulateCheckboxNote(noteID)

	}()

	return js.Null()
}

//DeleteCheckboxNoteItem DeleteCheckboxNoteItem
func (n *NoteHandler) DeleteCheckboxNoteItem(this js.Value, args []js.Value) interface{} {

	fmt.Println("id", args[0])
	fmt.Println("noteId", args[1])
	//fmt.Println("checked", args[2])
	//fmt.Println("text", args[3])

	var idInt = args[0].Int()
	var id = int64(idInt)
	//var idStr = strconv.FormatInt(id, 10)
	var noteIDInt = args[1].Int()
	var noteID = int64(noteIDInt)

	go func() {

		res := n.API.DeleteCheckboxItem(id)
		fmt.Println("cb delete suc: ", res.Success)

		n.pupulateCheckboxNote(noteID)

	}()

	return js.Null()
}

//UpdateCheckboxNoteTitle UpdateCheckboxNoteTitle
func (n *NoteHandler) UpdateCheckboxNoteTitle(this js.Value, args []js.Value) interface{} {
	document := js.Global().Get("document")
	title := document.Call("getElementById", "checkboxTitle").Get("value").String()
	fmt.Println("checkbox note title", title)
	noteIDStr := document.Call("getElementById", "noteId").Get("value").String()
	fmt.Println("checkbox note noteIDStr", noteIDStr)
	noteID, _ := strconv.ParseInt(noteIDStr, 10, 64)
	fmt.Println("checkbox note noteID", noteID)
	go func() {
		cbn := n.API.GetCheckboxNote(noteID)
		var unote api.Note
		unote.ID = noteID
		unote.Title = title
		unote.OwnerEmail = cbn.OwnerEmail
		unote.Type = cbn.Type
		res := n.API.UpdateNote(&unote)
		fmt.Println("cb note update suc: ", res.Success)

		n.pupulateCheckboxNote(noteID)

	}()

	return js.Null()
}

func (n *NoteHandler) pupulateCheckboxNote(noteID int64) {

	note := n.API.GetCheckboxNote(noteID)
	fmt.Println("checkbox note", *note)
	document := js.Global().Get("document")
	// document.Call("getElementById", "noteList").Get("style").Call("setProperty", "display", "none")
	//document.Call("getElementById", "checkboxNote").Get("style").Call("setProperty", "display", "block")
	document.Call("getElementById", "checkboxTitle").Set("value", note.Title)
	var rowHTML = ""
	var niddStr = strconv.FormatInt(noteID, 10)
	for _, nt := range note.NoteItems {
		fmt.Println("checkbox", nt)
		var idStr = strconv.FormatInt(nt.ID, 10)
		var nidStr = strconv.FormatInt(nt.NoteID, 10)
		var ched = ""

		if nt.Checked {
			ched = "checked"
		}
		// rowHTML = rowHTML + "<div class='form-group form-check'>"
		// rowHTML = rowHTML + "<input onchange='updateCheckItem(" + idStr + "," + nidStr + "," + "\"" + ched + "\"," + "\"" + nt.Text + "\"" + ")' type='checkbox' class='form-check-input' " + ched + ">"
		// rowHTML = rowHTML + "<input id='" + idStr + "' onchange='updateCheckItem(" + idStr + "," + nidStr + "," + "\"" + ched + "\"," + "\"" + nt.Text + "\"" + ")' type='text' class='form-control' value=\"" + nt.Text + "\"" + ">"
		// rowHTML = rowHTML + "</div>"
		rowHTML = rowHTML + "<div class='form-row'>"
		rowHTML = rowHTML + "<div class='form-group form-row-l'>"
		rowHTML = rowHTML + "<div class='form-check'>"
		rowHTML = rowHTML + "<input onchange='updateCheckItem(" + idStr + "," + nidStr + "," + "\"" + ched + "\"," + "\"" + nt.Text + "\"" + ")' type='checkbox' class='form-check-input' " + ched + ">"
		rowHTML = rowHTML + "<input id='" + idStr + "' onchange='updateCheckItem(" + idStr + "," + nidStr + "," + "\"" + ched + "\"," + "\"" + nt.Text + "\"" + ")' type='text' class='form-control' value=\"" + nt.Text + "\"" + ">"
		rowHTML = rowHTML + "</div>"
		rowHTML = rowHTML + "</div>"
		rowHTML = rowHTML + "<div class='form-group form-row-r'>"
		rowHTML = rowHTML + "<button onclick='deleteCheckItem(" + idStr + "," + nidStr + ")' type='button' class='btn btn-danger delete-btn'>X</button>"
		rowHTML = rowHTML + "</div>"
		rowHTML = rowHTML + "</div>"

	}
	rowHTML = rowHTML + "<div class='form-group form-check'>"
	//rowHTML = rowHTML + "<input type='checkbox' class='form-check-input'>"
	rowHTML = rowHTML + "<input id='newtxt' onchange='addCheckItem(" + niddStr + ")' type='text' class='form-control' style='width: 80%; margin: 0 0 0 2%;' placeholder='Add another item'>"
	rowHTML = rowHTML + "</div>"
	fmt.Println("rowHTML: ", rowHTML)
	document.Call("getElementById", "checkboxes").Set("innerHTML", rowHTML)
}
