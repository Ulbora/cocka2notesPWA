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
	//api "github.com/Ulbora/cocka2notesApi"
	//"strconv"
	"syscall/js"
)

//UpdateCheckboxNoteItem UpdateCheckboxNoteItem
func (n *NoteHandler) UpdateCheckboxNoteItem(this js.Value, args []js.Value) interface{} {

	fmt.Println("id", args[0])
	fmt.Println("noteId", args[1])
	fmt.Println("checked", args[2])
	fmt.Println("text", args[3])
	return js.Null()
}
