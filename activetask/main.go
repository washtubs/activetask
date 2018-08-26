package main

import "github.com/washtubs/activetask"
import "os"

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "notify" {
			activetask.Notify()
			os.Exit(0)
		}
	}
	activetask.Start()
}
