package server

import "github.com/fatih/structs"

func SetProjectGlobals() {
	structs.DefaultTagName = "json"
}
