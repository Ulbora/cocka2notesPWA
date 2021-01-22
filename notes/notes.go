package notes

import (
	"fmt"
	lg "github.com/Ulbora/Level_Logger"
	api "github.com/Ulbora/cocka2notesApi"
	"syscall/js"
)

//Handler Handler
type Handler interface {
	Login(this js.Value, args []js.Value) interface{}
	GetNoteList(this js.Value, args []js.Value) interface{}
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
	// // cookie1 := js.Global().Get("document").Get("username").String()
	// // fmt.Println("cookie1: ", cookie1)
	// fmt.Println("n.API", n.API)
	// // noteList := n.API.GetUsersNotes("tester@tester.com")
	// // fmt.Println(*noteList)
	// document := js.Global().Get("document")
	// document.Call("getElementById", "job").Set("value", "Engineer of Notes to go")
	// return js.Null()
	go func() {
		noteList := n.API.GetUsersNotes("tester@tester.com")
		fmt.Println(*noteList)
		document := js.Global().Get("document")
		document.Call("getElementById", "job").Set("value", (*noteList)[0].Title)
		//email := js.Global().Get("document").Get("email").String()
		//fmt.Println(email)
	}()
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
		// user := n.API.GetUser(newu.Text)
		// fmt.Println(user)

		// document.Call("getElementById", "job").Set("value", (*noteList)[0].Title)
		//email := js.Global().Get("document").Get("email").String()
		//fmt.Println(email)
	}()
	return js.Null()
}
