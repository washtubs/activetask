package activetask

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func IsNotWorking(taskMessage string) bool {
	return taskMessage == ""
}

// watches the activetask file for new messages
func watchTask(newTask chan string) {
	last := "" // force the channel to always fire in the first cycle

	// Really important that we don't use `time.Tick` every 2s as that causes a resource leak of one goroutine every 2s.
	// In practice, this starts causing performance degradation (>100% CPU) after several days
	// I see ~180% CPU usage for a process that's been running for 713h,
	// which goes to show that it's very mild but it definitely adds up.
	ticker := time.NewTicker(time.Second * 2)
	for {
		<-ticker.C

		current := getTaskMessage()

		if current != last {
			newTask <- current
			last = current
			log.Printf("task %s", current)
		}
	}
}

func watchNotification(notificationRequest chan bool) {
	ticker := time.NewTicker(time.Second * 2)
	for {
		<-ticker.C
		if GetNotifyRequest() {
			RemoveNotifyRequest()
			notificationRequest <- true
		}
	}
}

func Notify() {
	PutNotifyRequest()
}

func Start() {
	newTaskChan := make(chan string, 1)
	go watchTask(newTaskChan)

	manualReminder := make(chan bool)
	go watchNotification(manualReminder)

	var cancelReminders chan bool
	for {
		// got a task, but might be -1
		task := <-newTaskChan

		// asyncronously ensure that the last IssueReminders is cleaned up
		go func(cancel chan bool) {
			select {
			case cancel <- true:
			default:
			}
		}(cancelReminders)

		// asynchronously set up reminders, with a channel that ensures cancellation
		cancelReminders = make(chan bool)
		go IssueRemindersAndLogTime(time.Now(), task, manualReminder, cancelReminders)
	}

}

func Watch(includeNotWorking bool, command string) error {
	newTaskChan := make(chan string, 1)
	go watchTask(newTaskChan)
	for {
		// got a task, but might be empty
		task := <-newTaskChan

		if includeNotWorking || task != "" {
			cmd := exec.Command(command, task)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if err != nil {
				log.Printf("Got error executing command=[%+v]: %s", cmd, err.Error())
				return err
			}

		}
	}

}

func GetTaskMessage() string {
	task := getTaskMessage()
	if task == "" {
		return "No task"
	} else {
		return fmt.Sprintf("%s", task)
	}
}
