package notes

import (
	"fmt"
	//lg "github.com/Ulbora/Level_Logger"
	//api "github.com/Ulbora/cocka2notesApi"
	"strconv"
	"syscall/js"
)

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
		} else if args[1].String() == "checkbox" {
			note := n.API.GetCheckboxNote(id)
			fmt.Println("checkbox note", *note)
			document := js.Global().Get("document")
			document.Call("getElementById", "noteList").Get("style").Call("setProperty", "display", "none")
			document.Call("getElementById", "checkboxNote").Get("style").Call("setProperty", "display", "block")
			document.Call("getElementById", "checkboxTitle").Set("value", note.Title)
			var rowHTML = ""
			for _, nt := range note.NoteItems {
				fmt.Println("checkbox", nt)
				var idStr = strconv.FormatInt(nt.ID, 10)
				var nidStr = strconv.FormatInt(nt.NoteID, 10)
				var ched = ""

				if nt.Checked {
					ched = "checked"
				}
				rowHTML = rowHTML + "<div class='form-group form-check'>"
				rowHTML = rowHTML + "<input onchange='updateCheckItem(" + idStr + "," + nidStr + "," + "\"" + ched + "\"," + "\"" + nt.Text + "\"" + ")' type='checkbox' class='form-check-input' " + ched + ">"
				rowHTML = rowHTML + "<input type='text' class='form-control' value=\"" + nt.Text + "\"" + ">"
				rowHTML = rowHTML + "</div>"

			}
			fmt.Println("rowHTML: ", rowHTML)
			document.Call("getElementById", "checkboxes").Set("innerHTML", rowHTML)
		}

	}()

	//noteList := n.API.GetNote(email)
	return js.Null()
}
