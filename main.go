package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

//Config Config
type Config struct {
	URL    string `json:"url"`
	APIKey string `json:"apiKey"`
}

func main() {
	fmt.Println("nh.API start of main:", nh.API)

	wg.Add(1)
	// var nh ns.NoteHandler
	var napi api.NotesAPI
	var head api.Headers
	napi.SetHeader(&head)
	//napi.SetRestURL("http://localhost:3000")
	//napi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	nh.API = napi.GetNew()

	//api := napi.GetNew()
	var l lg.Logger
	napi.SetLogger(&l)
	napi.SetLogLevel(lg.AllLevel)

	h := nh.GetNew()

	fmt.Println("nh.API", nh.API)
	fmt.Println("Hello, WebAssembly, how you doing!")

	js.Global().Set("getNotes", js.FuncOf(h.GetNoteList))
	js.Global().Set("showNote", js.FuncOf(h.GetNote))
	js.Global().Set("addNewNote", js.FuncOf(h.AddNote))
	js.Global().Set("addNote", js.FuncOf(h.AddNewNote))
	js.Global().Set("deleteNote", js.FuncOf(h.DeleteNote))

	js.Global().Set("updateCheckTitle", js.FuncOf(h.UpdateCheckboxNoteTitle))
	js.Global().Set("updateCheckItem", js.FuncOf(h.UpdateCheckboxNoteItem))
	js.Global().Set("addCheckItem", js.FuncOf(h.AddCheckboxNoteItem))
	js.Global().Set("deleteCheckItem", js.FuncOf(h.DeleteCheckboxNoteItem))

	js.Global().Set("updateTextTitle", js.FuncOf(h.UpdateTextNoteTitle))
	js.Global().Set("updateTextItem", js.FuncOf(h.UpdateTextNoteItem))
	js.Global().Set("addTextItem", js.FuncOf(h.AddTextNoteItem))
	js.Global().Set("deleteTextItem", js.FuncOf(h.DeleteTextNoteItem))

	js.Global().Set("addUsersToNote", js.FuncOf(h.AddUsersToNote))
	js.Global().Set("addNoteUser", js.FuncOf(h.AddNoteUser))

	js.Global().Set("login", js.FuncOf(h.Login))
	js.Global().Set("loginScreen", js.FuncOf(h.LoginScreen))
	js.Global().Set("changePwScreen", js.FuncOf(h.ChangePwScreen))
	js.Global().Set("changePassword", js.FuncOf(h.ChangePassword))
	js.Global().Set("resetPwScreen", js.FuncOf(h.ResetPwScreen))
	js.Global().Set("resetPassword", js.FuncOf(h.ResetPassword))

	js.Global().Set("registerScreen", js.FuncOf(h.RegisterScreen))
	js.Global().Set("register", js.FuncOf(h.Register))

	// js.Global().Set("getNotes", MyGoFunc)

	go func() {
		req, rErr := http.NewRequest("GET", "http://www.cocka2notes.com/rs/config/get", nil)
		fmt.Println("req err: ", rErr)
		client := &http.Client{}
		resp, err := client.Do(req)
		var conf Config
		if err != nil {
			fmt.Println("Do err: ", err)
			fmt.Println("resp in fail: ", resp)
		} else {
			defer resp.Body.Close()
			status := resp.StatusCode
			fmt.Println("status: ", status)
			//fmt.Println("resp.Body: ", resp.Body)
			//fmt.Println("resp err: ", err)
			decoder := json.NewDecoder(resp.Body)
			error := decoder.Decode(&conf)
			if error != nil {
				fmt.Println("Decode Error:  ", error.Error())
			}
		}
		fmt.Println("conf: ", conf)
		fmt.Println("url: ", conf.URL)
		fmt.Println("api: ", conf.APIKey)
		if conf.APIKey != "" && conf.URL != "" {
			napi.SetRestURL(conf.URL)
			napi.SetAPIKey(conf.APIKey)
		} else {
			napi.SetRestURL("http://localhost:3000")
			napi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
		}
		emailFn := js.Global().Get("getUserEmail")
		cemail := emailFn.Invoke()
		fmt.Println("email: ", cemail)
		if cemail.String() == "" {
			fmt.Println("email: ", cemail)
			document := js.Global().Get("document")
			document.Call("getElementById", "loginScreen").Get("style").Call("setProperty", "display", "block")
		} else {
			//nh.Email = cemail.String()
			nh.PopulateNoteList(cemail.String())
		}
		//else  go go note list
	}()

	wg.Wait()
	//cookies := js.Global().Get("document").Get("cookie").String()
}

// go mod init github.com/Ulbora/cocka2notesWA
