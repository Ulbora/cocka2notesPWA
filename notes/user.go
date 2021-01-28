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
	//"strconv"
	"syscall/js"
)

//LoginScreen LoginScreen
func (n *NoteHandler) LoginScreen(this js.Value, args []js.Value) interface{} {
	document := js.Global().Get("document")
	document.Call("getElementById", "loginScreen").Get("style").Call("setProperty", "display", "block")
	document.Call("getElementById", "noteList").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "textNote").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "checkboxNote").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "noteUserForm").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "noteUserTable").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "addNoteForm").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "changePwScreen").Get("style").Call("setProperty", "display", "none")
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
		var u api.User
		u.Email = email
		u.Password = pw
		res := n.API.Login(&u)
		fmt.Println("login suc: ", *res)
		if res.Success {
			emailc := js.Global().Get("setUserEmail")
			emailc.Invoke(email)
			//n.Email = email
			document := js.Global().Get("document")
			document.Call("getElementById", "loginScreen").Get("style").Call("setProperty", "display", "none")
			//go to note list
			n.PopulateNoteList(email)
		} else if res.Email == "" {
			// auto create a new account
			suc := n.API.AddUser(&u)
			if suc.Success {
				emailc := js.Global().Get("setUserEmail")
				emailc.Invoke(email)
				document := js.Global().Get("document")
				document.Call("getElementById", "loginScreen").Get("style").Call("setProperty", "display", "none")
				// gto to note list
				n.PopulateNoteList(email)
			}
		}

	}()
	return js.Null()
}

//ChangePwScreen ChangePwScreen
func (n *NoteHandler) ChangePwScreen(this js.Value, args []js.Value) interface{} {
	document := js.Global().Get("document")
	document.Call("getElementById", "loginScreen").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "pwnotMatch").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "changePwScreen").Get("style").Call("setProperty", "display", "block")
	emailFn := js.Global().Get("getUserEmail")
	cemail := emailFn.Invoke()
	document.Call("getElementById", "cpwemail").Set("value", cemail.String())
	document.Call("getElementById", "newPassword").Set("value", "")
	document.Call("getElementById", "newPassword2").Set("value", "")
	//n.PopulateNoteList(cemail.String())
	return js.Null()
}

//ChangePassword ChangePassword
func (n *NoteHandler) ChangePassword(this js.Value, args []js.Value) interface{} {
	document := js.Global().Get("document")
	npw := document.Call("getElementById", "newPassword").Get("value").String()
	npw2 := document.Call("getElementById", "newPassword2").Get("value").String()
	if npw != "" && npw == npw2 {
		emailFn := js.Global().Get("getUserEmail")
		cemail := emailFn.Invoke()
		var u api.User
		u.Email = cemail.String()
		u.Password = npw
		go func() {
			res := n.API.UpdateUser(&u)
			if res.Success {
				n.PopulateNoteList(cemail.String())
			}
		}()
	} else {
		document.Call("getElementById", "pwnotMatch").Get("style").Call("setProperty", "display", "block")
	}
	return js.Null()
}
