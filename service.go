package activetask

import "log"

import "time"

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
	last := -999 // force the channel to always fire immediately
	for {
		<-time.Tick(time.Second * 2)
		current := GetTaskId()
		if IsNotWorking(current) {
			current = notWorkingTaskId
		}
		if current != last {
			newTaskId <- current
			last = current
			log.Printf("id %d", current)
		}
	}
}

func watchNotification(notificationRequest chan bool) {
	for {
		<-time.Tick(time.Second * 2)
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
	newTaskChan := make(chan int)
	go watchTaskId(newTaskChan)

	manualReminder := make(chan bool)
	go watchNotification(manualReminder)

	var cancelReminders chan bool
	for {
		// got a task, but might be -1
		current := <-newTaskChan

		// asyncronously ensure that the last IssueReminders is cleaned up
		go func(cancel chan bool) { cancel <- true }(cancelReminders)

		task := GetTaskById(current)
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
