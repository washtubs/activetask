package activetask

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func OpenActiveTaskFileOrPanic() *os.File {

	activetaskPath := os.Getenv("HOME") + "/.activetask"

	f, err := os.Open(activetaskPath)
	if err != nil {
		if os.IsNotExist(err) {
			f, err := os.Create(activetaskPath)
			if err != nil {
				panic("Failed to create " + activetaskPath + " : " + err.Error())
			}
			log.Println(activetaskPath + " created.")
			return f
		}
		panic("~/.activetask file inaccessible : " + err.Error())
	}
	return f

}

func GetTaskId() int {
	f := OpenActiveTaskFileOrPanic()
	defer f.Close()
	scanner := bufio.NewScanner(f)
	if scanner.Scan() {
		answer, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return -1
		} else {
			return answer
		}
	} else {
		return -1
	}
}
