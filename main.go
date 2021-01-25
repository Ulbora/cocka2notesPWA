package main

import (
	"fmt"
	lg "github.com/Ulbora/Level_Logger"
	api "github.com/Ulbora/cocka2notesApi"
	ns "github.com/Ulbora/cocka2notesWA/notes"
	//"os"
	"sync"
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

var wg sync.WaitGroup
var nh ns.NoteHandler

func main() {

	//email := js.Global().Get("document").Get("email").String()
	//fmt.Println("email: ", email)
	// if cookie1 != "ben" {
	//js.Global().Get("document").Set("email", "ben")
	// }

	// mailHost := os.Getenv("EMAIL_HOST")
	// fmt.Println("email host: ", mailHost)
	wg.Add(1)
	// var nh ns.NoteHandler
	var napi api.NotesAPI
	var head api.Headers
	napi.SetHeader(&head)
	napi.SetRestURL("http://localhost:3000")
	napi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	nh.API = napi.GetNew()

	//api := napi.GetNew()
	var l lg.Logger
	napi.SetLogger(&l)
	//napi.SetLogLevel(lg.AllLevel)

	h := nh.GetNew()

	// noteList := api.GetUsersNotes("tester@tester.com")
	// noteList := nh.API.GetUsersNotes("tester@tester.com")
	fmt.Println("nh.API", nh.API)
	fmt.Println("Hello, WebAssembly, how you doing!")
	// fmt.Println(*noteList)
	// document := js.Global().Get("document")
	// p := document.Call("createElement", "p")
	// p.Set("innerHTML", "Hello WASM from Go!")
	// document.Get("body").Call("appendChild", p)
	// document.Call("getElementById", "name").Set("value", "ken")

	//hello := js.Global().Get("sayHello")
	//hello.Invoke()

	//something := js.Global().Get("saySomething")
	//something.Invoke("How you doing?")

	// js.Global().Set("getNotes", js.FuncOf(getNoteList))

	js.Global().Set("getNotes", js.FuncOf(h.GetNoteList))
	js.Global().Set("showNote", js.FuncOf(h.GetNote))
	js.Global().Set("updateCheckTitle", js.FuncOf(h.UpdateCheckboxNoteTitle))
	js.Global().Set("updateCheckItem", js.FuncOf(h.UpdateCheckboxNoteItem))
	js.Global().Set("addCheckItem", js.FuncOf(h.AddCheckboxNoteItem))
	js.Global().Set("deleteCheckItem", js.FuncOf(h.DeleteCheckboxNoteItem))
	js.Global().Set("login", js.FuncOf(h.Login))

	//func Clone() js.Func {
	// cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// 	noteList := nh.API.GetUsersNotes("tester@tester.com")
	// 	//fmt.Println(*noteList)
	// 	document := js.Global().Get("document")
	// 	document.Call("getElementById", "job").Set("value", "Engineer of Notes to go")
	// 	document.Call("getElementById", "job").Set("value", *noteList)
	// 	cb.Release() // release the function if the button will not be clicked again
	// 	return js.Null()
	// })

	// js.Global().Set("getNotes", MyGoFunc)

	go func() {
		emailFn := js.Global().Get("getUserEmail")
		cemail := emailFn.Invoke()
		fmt.Println("email: ", cemail)
		if cemail.String() == "" {
			fmt.Println("email: ", cemail)
			document := js.Global().Get("document")
			document.Call("getElementById", "loginScreen").Get("style").Call("setProperty", "display", "block")
		} else {
			nh.Email = cemail.String()
			nh.PopulateNoteList(cemail.String())
		}
		//else  go go note list
	}()

	wg.Wait()
	//cookies := js.Global().Get("document").Get("cookie").String()
}

// func getNoteList(this js.Value, args []js.Value) interface{} {
// 	go func() {
// 		noteList := nh.API.GetUsersNotes("tester@tester.com")
// 		fmt.Println(*noteList)
// 		document := js.Global().Get("document")
// 		document.Call("getElementById", "job").Set("value", (*noteList)[0].Title)
// 		email := js.Global().Get("document").Get("email").String()
// 		fmt.Println(email)
// 	}()
// 	return js.Null()
// }

// go mod init github.com/Ulbora/cocka2notesWA
