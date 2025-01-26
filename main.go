/*
	NautilusDB is database that has been created to understand how databases work under the hood
    Copyright (C) 2025  Snigdhadeb Roy Chowdhury

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.

	For communications or inquiries, contact:
	Email: snigdhadeb_roychowdhury@outlook.com
*/

package main

import (
	"NautilusDB/location"
	"fmt"
)

func main() {
	fmt.Println("Hello, this NautilusDB")
	// pageSize := syscall.Getpagesize()
	// fmt.Printf("Page size: %d bytes\n", pageSize)
	// cmd.Execute()
	err := location.CheckLocation("employees")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
