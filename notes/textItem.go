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
	"encoding/json"
	"fmt"

	//lg "github.com/Ulbora/Level_Logger"
	"strconv"
	"syscall/js"

	api "github.com/Ulbora/cocka2notesApi"
)

//UpdateTextNoteTitle UpdateTextNoteTitle
func (n *NoteHandler) UpdateTextNoteTitle(this js.Value, args []js.Value) interface{} {
	document := js.Global().Get("document")
	title := document.Call("getElementById", "textTitle").Get("value").String()
	fmt.Println("text note title", title)
	noteIDStr := document.Call("getElementById", "noteId").Get("value").String()
	fmt.Println("text note noteIDStr", noteIDStr)
	noteID, _ := strconv.ParseInt(noteIDStr, 10, 64)
	fmt.Println("text note noteID", noteID)
	go func() {
		cbn := n.API.GetNote(noteID)
		var unote api.Note
		unote.ID = noteID
		unote.Title = title
		unote.OwnerEmail = cbn.OwnerEmail
		unote.Type = cbn.Type
		res := n.API.UpdateNote(&unote)
		fmt.Println("cb note update suc: ", res.Success)

		nlst := n.API.GetNoteList()
		JSON, _ := json.Marshal(nlst)
		n.SaveToLocalStorage("noteList", JSON)

		n.PopulateTextNote(noteID)

	}()

	return js.Null()
}

//UpdateTextNoteItem UpdateTextNoteItem
func (n *NoteHandler) UpdateTextNoteItem(this js.Value, args []js.Value) interface{} {
	fmt.Println("id", args[0])
	fmt.Println("noteId", args[1])
	fmt.Println("text", args[2])
	var idInt = args[0].Int()
	var id = int64(idInt)
	var idStr = strconv.FormatInt(id, 10)
	var noteIDInt = args[1].Int()
	var noteID = int64(noteIDInt)
	document := js.Global().Get("document")

	go func() {
		txt := document.Call("getElementById", idStr).Get("value").String()
		//title := document.Call("getElementById", "textTitle").Get("value").String()
		fmt.Println("text", txt)
		n.API.FlushFailedCache()
		//cbn := n.API.GetNote(noteID)
		var item api.NoteItem
		item.ID = id
		item.NoteID = noteID
		item.Text = txt
		res := n.API.UpdateNoteItem(&item)
		fmt.Println("textnote item update suc: ", res.Success)

		nlst := n.API.GetNoteList()
		JSON, _ := json.Marshal(nlst)
		n.SaveToLocalStorage("noteList", JSON)

		n.PopulateTextNote(noteID)

	}()

	return js.Null()
}

//AddTextNoteItem AddTextNoteItem
func (n *NoteHandler) AddTextNoteItem(this js.Value, args []js.Value) interface{} {

	fmt.Println("noteId", args[0])
	var noteIDInt = args[0].Int()
	var noteID = int64(noteIDInt)
	fmt.Println("noteId", noteID)
	//var idInt = args[0].Int()
	document := js.Global().Get("document")
	txt := document.Call("getElementById", "tnewtxt").Get("value").String()
	fmt.Println("tnewtxt", txt)
	go func() {
		n.API.FlushFailedCache()
		// txt := document.Call("getElementById", "tnewtxt").Get("value").String()
		// fmt.Println("txt", txt)
		// txt = document.Call("getElementById", "tnewtxt").Get("value").String()
		// fmt.Println("txt", txt)
		var item api.NoteItem
		item.NoteID = noteID
		item.Text = txt
		res := n.API.AddNoteItem(&item)
		fmt.Println("cb update suc: ", res.Success)

		nlst := n.API.GetNoteList()
		JSON, _ := json.Marshal(nlst)
		n.SaveToLocalStorage("noteList", JSON)

		n.PopulateTextNote(noteID)

	}()
	return js.Null()
}

//DeleteTextNoteItem DeleteTextNoteItem
func (n *NoteHandler) DeleteTextNoteItem(this js.Value, args []js.Value) interface{} {
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
		res := n.API.DeleteNoteItem(id)
		fmt.Println("cb delete suc: ", res.Success)
		n.PopulateTextNote(noteID)

	}()

	return js.Null()
}

func (n *NoteHandler) PopulateTextNote(noteID int64) {
	n.API.FlushFailedCache()
	note := n.API.GetNote(noteID)
	fmt.Println("textbox note", *note)
	document := js.Global().Get("document")
	document.Call("getElementById", "noteList").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "textNote").Get("style").Call("setProperty", "display", "block")
	document.Call("getElementById", "textTitle").Set("value", note.Title)
	//document.Call("getElementById", "noteId").Set("value", idInt)
	//document.Call("getElementById", "noteId").Set("value", noteIDStr)
	var rowHTML = ""
	var niddStr = strconv.FormatInt(noteID, 10)
	for _, nt := range note.NoteItems {
		fmt.Println("checkbox", nt)
		var idStr = strconv.FormatInt(nt.ID, 10)
		var nidStr = strconv.FormatInt(nt.NoteID, 10)
		//var ched = ""

		// if nt.Checked {
		// 	ched = "checked"
		// }
		rowHTML = rowHTML + "<div class='form-row'>"
		rowHTML = rowHTML + "<div class='form-group form-row-l'>"
		//rowHTML = rowHTML + "<div class='form-check'>"
		//rowHTML = rowHTML + "<input onchange='updateCheckItem(" + idStr + "," + nidStr + "," + "\"" + ched + "\"," + "\"" + nt.Text + "\"" + ")' type='checkbox' class='form-check-input' " + ched + ">"
		rowHTML = rowHTML + "<input id='" + idStr + "' onchange='updateTextItem(" + idStr + "," + nidStr + "," + "\"" + nt.Text + "\"" + ")' type='text' class='form-control' value=\"" + nt.Text + "\"" + ">"
		rowHTML = rowHTML + "</div>"
		//rowHTML = rowHTML + "</div>"
		rowHTML = rowHTML + "<div class='form-group form-row-r'>"
		rowHTML = rowHTML + "<button onclick='deleteTextItem(" + idStr + "," + nidStr + ")' type='button' class='btn btn-danger delete-btn'>X</button>"
		rowHTML = rowHTML + "</div>"
		rowHTML = rowHTML + "</div>"

	}
	rowHTML = rowHTML + "<div class='form-group'>"
	//rowHTML = rowHTML + "<input type='checkbox' class='form-check-input'>"
	rowHTML = rowHTML + "<input id='tnewtxt' onchange='addTextItem(" + niddStr + ")' type='text' class='form-control' style='width: 80%; margin: 0 0 0 2%;' placeholder='Add another item'>"
	rowHTML = rowHTML + "</div>"

	fmt.Println("rowHTML: ", rowHTML)
	document.Call("getElementById", "items").Set("innerHTML", rowHTML)
}
