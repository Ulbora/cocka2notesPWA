package notes

import (
	"fmt"
	//lg "github.com/Ulbora/Level_Logger"
	//api "github.com/Ulbora/cocka2notesApi"
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
			note := n.API.GetNote(id)
			fmt.Println("note", *note)
			document := js.Global().Get("document")
			document.Call("getElementById", "noteList").Get("style").Call("setProperty", "display", "none")
			document.Call("getElementById", "textNote").Get("style").Call("setProperty", "display", "block")
			document.Call("getElementById", "textTitle").Set("value", note.Title)
			document.Call("getElementById", "noteId").Set("value", idInt)
			var rowHTML = ""
			var niddStr = strconv.FormatInt(id, 10)
			for _, nt := range note.NoteItems {
				fmt.Println("textItem", nt)
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
				rowHTML = rowHTML + "<button onclick='deleteNoteItem(" + idStr + "," + nidStr + ")' type='button' class='btn btn-danger delete-btn'>X</button>"
				rowHTML = rowHTML + "</div>"
				rowHTML = rowHTML + "</div>"

			}
			rowHTML = rowHTML + "<div class='form-group'>"
			//rowHTML = rowHTML + "<input type='checkbox' class='form-check-input'>"
			rowHTML = rowHTML + "<input id='newtxt' onchange='addTextItem(" + niddStr + ")' type='text' class='form-control'>"
			rowHTML = rowHTML + "</div>"

			fmt.Println("rowHTML: ", rowHTML)
			document.Call("getElementById", "items").Set("innerHTML", rowHTML)
		} else if args[1].String() == "checkbox" {
			note := n.API.GetCheckboxNote(id)
			fmt.Println("checkbox note", *note)
			document := js.Global().Get("document")
			document.Call("getElementById", "noteList").Get("style").Call("setProperty", "display", "none")
			document.Call("getElementById", "checkboxNote").Get("style").Call("setProperty", "display", "block")
			document.Call("getElementById", "checkboxTitle").Set("value", note.Title)
			document.Call("getElementById", "noteId").Set("value", idInt)
			var rowHTML = ""
			var niddStr = strconv.FormatInt(id, 10)
			for _, nt := range note.NoteItems {
				fmt.Println("checkbox", nt)
				var idStr = strconv.FormatInt(nt.ID, 10)
				var nidStr = strconv.FormatInt(nt.NoteID, 10)
				var ched = ""

				if nt.Checked {
					ched = "checked"
				}
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
			rowHTML = rowHTML + "<input id='newtxt' onchange='addCheckItem(" + niddStr + ")' type='text' class='form-control'>"
			rowHTML = rowHTML + "</div>"

			fmt.Println("rowHTML: ", rowHTML)
			document.Call("getElementById", "checkboxes").Set("innerHTML", rowHTML)
		}

	}()

	//noteList := n.API.GetNote(email)
	return js.Null()
}
