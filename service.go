package activetask

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"
)

const notWorkingTaskId = -1

func IsNotWorking(taskId int) bool {
	if taskId < 1 {
		// IDs below 1 are impossible
		return true
	}

	// actually look up the task
	task := GetTaskById(taskId)
	if task == nil {
		// if it doesn't exist it's deleted or something.
		return true
	} else if task.Completed {
		// if it's completed we're not actually working on anything
		return true
	}

	return false
}

// watches the activetask file
// if the ID corresponds to an incomplete task, place that ID on the channel
// if note, place the `notWorkingTaskId` on the channel
func watchTaskId(newTaskId chan int) {
	last := -999 // force the channel to always fire in the first cycle

	// Really important that we don't use `time.Tick` every 2s as that causes a resource leak of one goroutine every 2s.
	// In practice, this starts causing performance degradation (>100% CPU) after several days
	// I see ~180% CPU usage for a process that's been running for 713h,
	// which goes to show that it's very mild but it definitely adds up.
	ticker := time.NewTicker(time.Second * 2)
	for {
		<-ticker.C

		currentId := GetTaskId()
		if IsNotWorking(currentId) {
			currentId = notWorkingTaskId
		}

		if currentId != last {
			newTaskId <- currentId
			last = currentId
			log.Printf("id %d", currentId)
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
	newTaskChan := make(chan int, 1)
	go watchTaskId(newTaskChan)

	manualReminder := make(chan bool)
	go watchNotification(manualReminder)

	var cancelReminders chan bool
	for {
		// got a task, but might be -1
		currentId := <-newTaskChan

		// asyncronously ensure that the last IssueReminders is cleaned up
		go func(cancel chan bool) {
			select {
			case cancel <- true:
			default:
			}
		}(cancelReminders)

		task := GetTaskById(currentId)
		if task != nil {
			log.Print("Got task " + task.Subject)
		} else {
			log.Print("Got no task.")
		}

		// asynchronously set up reminders, with a channel that ensures cancellation
		cancelReminders = make(chan bool)
		go IssueRemindersAndLogTime(time.Now(), task, manualReminder, cancelReminders)
	}

}

func Watch(includeNotWorking bool, shellCommand string) error {
	newTaskChan := make(chan int, 1)
	go watchTaskId(newTaskChan)
	for {
		// got a task, but might be -1
		currentId := <-newTaskChan

		task := GetTaskById(currentId)

		if includeNotWorking || task != nil {
			t, err := template.New(shellCommand).Funcs(template.FuncMap{"StringsJoin": strings.Join}).Parse(shellCommand)
			if err != nil {
				return err
			}

			buf := bytes.NewBufferString("")
			err = t.Execute(buf, task)
			if err != nil {
				log.Printf("Failed to execute template with the given task %+v: %s", task, err.Error())
				return err
			}

			shellCommandRendered := buf.String()
			cmd := exec.Command("sh", "-c", shellCommandRendered)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			err = cmd.Run()
			if err != nil {
				log.Printf("Got error executing command=[%s]: %s", shellCommandRendered, err.Error())
				return err
			}

		}
	}

}

func GetTaskMessage() string {
	taskId := GetTaskId()
	task := GetTaskById(taskId)
	if task == nil || task.Completed {
		return "No task"
	} else {
		return fmt.Sprintf("#%d %s", taskId, task.Subject)
	}
}
