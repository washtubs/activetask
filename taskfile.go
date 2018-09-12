package activetask

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func openFileOrPanic(file string, readOnly bool) *os.File {

	var f *os.File
	var err error

	if readOnly {
		f, err = os.Open(file)
	} else {
		f, err = os.Create(file)
	}
	if err != nil {
		if os.IsNotExist(err) {
			f, err := os.Create(file)
			if err != nil {
				panic("Failed to create " + file + " : " + err.Error())
			}
			log.Println(file + " created.")
			return f
		}
		panic(file + " file inaccessible : " + err.Error())
	}
	return f

}

func getIdFromFile(file string) int {
	f := openFileOrPanic(file, true)
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

func GetTaskId() int {
	activetaskPath := os.Getenv("HOME") + "/.activetask"
	return getIdFromFile(activetaskPath)
}

func GetNotifyRequest() bool {
	notifyRequestPath := os.Getenv("HOME") + "/.activetask-notify"
	return getIdFromFile(notifyRequestPath) == 1
}

func PutNotifyRequest() {
	notifyRequestPath := filepath.Join(os.Getenv("HOME"), ".activetask-notify")
	f := openFileOrPanic(notifyRequestPath, false)
	defer f.Close()
	_, err := f.WriteString("1")
	if err != nil {
		panic(err.Error())
	}

}

func RemoveNotifyRequest() {
	notifyRequestPath := filepath.Join(os.Getenv("HOME"), ".activetask-notify")
	f := openFileOrPanic(notifyRequestPath, false)
	defer f.Close()
	_, err := f.WriteString("-1")
	if err != nil {
		panic(err.Error())
	}
}
