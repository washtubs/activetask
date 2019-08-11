package main

import (
	"fmt"
	"os"

	"github.com/washtubs/activetask"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "notify" {
			activetask.Notify()
		} else if os.Args[1] == "message" {
			fmt.Println(activetask.GetTaskMessage())
		}
		os.Exit(0)
	}
	activetask.Start()
}
