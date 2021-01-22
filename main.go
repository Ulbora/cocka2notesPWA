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

	// cookie1 := js.Global().Get("document").Get("username").String()
	// fmt.Println("cookie1: ", cookie1)
	// if cookie1 != "ben" {
	js.Global().Get("document").Set("username", "ben")
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

	// h := nh.GetNew()

	// noteList := api.GetUsersNotes("tester@tester.com")
	// noteList := nh.API.GetUsersNotes("tester@tester.com")
	fmt.Println("nh.API", nh.API)
	fmt.Println("Hello, WebAssembly, how you doing!")
	// fmt.Println(*noteList)
	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
	p.Set("innerHTML", "Hello WASM from Go!")
	document.Get("body").Call("appendChild", p)
	document.Call("getElementById", "name").Set("value", "ken")

	// js.Global().Set("getNotes", js.FuncOf(getNoteList))
	js.Global().Set("getNotes", js.FuncOf(getNoteList))
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

	wg.Wait()
	//cookies := js.Global().Get("document").Get("cookie").String()
}

func getNoteList(this js.Value, args []js.Value) interface{} {
	go func() {
		noteList := nh.API.GetUsersNotes("tester@tester.com")
		fmt.Println(*noteList)
		document := js.Global().Get("document")
		document.Call("getElementById", "job").Set("value", (*noteList)[0].Title)
		cookie1 := js.Global().Get("document").Get("username").String()
		fmt.Println(cookie1)
	}()
	return js.Null()
}

// go mod init github.com/Ulbora/cocka2notesWA
