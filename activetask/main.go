package main

import (
	"fmt"
	"log"
	"os"

	"github.com/washtubs/activetask"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "notify" {
			activetask.Notify()
		} else if os.Args[1] == "message" {
			fmt.Println(activetask.GetTaskMessage())
		} else if os.Args[1] == "ontask" {
			if len(os.Args) <= 2 {
				log.Fatal("ontask needs a shell command argument as the first argument")
			}
			err := activetask.Watch(false, os.Args[2])
			if err != nil {
				log.Fatal(err)
			}

		}
		os.Exit(0)
	}
	activetask.Start()
}
