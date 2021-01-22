package notes

import (
	"fmt"
	lg "github.com/Ulbora/Level_Logger"
	api "github.com/Ulbora/cocka2notesApi"
	"syscall/js"
)

type Handler interface {
	GetNoteList(this js.Value, args []js.Value) interface{}
}
type NoteHandler struct {
	API   api.API
	Log   *lg.Logger
	Email string
}

//GetNew GetNew
func (n *NoteHandler) GetNew() Handler {
	return n
}

func (n *NoteHandler) GetNoteList(this js.Value, args []js.Value) interface{} {
	// cookie1 := js.Global().Get("document").Get("username").String()
	// fmt.Println("cookie1: ", cookie1)
	fmt.Println("n.API", n.API)
	// noteList := n.API.GetUsersNotes("tester@tester.com")
	// fmt.Println(*noteList)
	document := js.Global().Get("document")
	document.Call("getElementById", "job").Set("value", "Engineer of Notes to go")
	return js.Null()
}
