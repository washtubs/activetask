package activetask

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gammons/todolist/todolist"
	"github.com/gen2brain/beeep"
)

// The idea behind a graduated interval is that sometimes we expect tasks to
// take a short time. And having the 5 minute reminder is nice to let you know
// that you might be taking longer than you anticipated.
// But obiously after a certain point you expect things to take a while, so
// faster reminders are just noise. Long tasks might take several hours. Short
// tasks might take less than 2 minutes.
var graduatedNotificationIntervals []time.Duration = []time.Duration{
	5 * time.Minute,
	10 * time.Minute,
	20 * time.Minute,
	40 * time.Minute,
	1 * time.Hour,
}

var graduatedNotificationIntervalsNoTask []time.Duration = []time.Duration{
	5 * time.Minute,
	10 * time.Minute,
	20 * time.Minute,
}

func LogTime(startTime time.Time, task *todolist.Todo) {
	if task != nil {
		beeep.Notify("[active-task] #"+strconv.Itoa(task.Id)+" "+task.Subject,
			"Logged "+humanize.RelTime(startTime, time.Now(), "", ""), "a")
	}
}

func fuzzDuration(duration time.Duration, fuzz float64) time.Duration {
	if rand.Intn(1) == 1 {
		fuzz = fuzz * -1
	}
	return duration + time.Duration(rand.Float64()*fuzz)
}

// Start issuing reminders on a graduated interval, indicating the time when the task started
// If the task is null, remind the user that they have no current task
func IssueRemindersAndLogTime(startTime time.Time, task *todolist.Todo, manualReminder chan bool, cancel chan bool) {
	var i, taskId int

	var intervals []time.Duration

	if task == nil {
		taskId = notWorkingTaskId
		intervals = graduatedNotificationIntervalsNoTask
	} else {
		taskId = task.Id
		intervals = graduatedNotificationIntervals
	}
	i = 0
	for {

		interval := fuzzDuration(intervals[i], 0.15)
		intervalTicked := false

		select {
		case <-cancel:
			log.Printf("Aborting reminders for %d", taskId)
			LogTime(startTime, task)
			return
		case <-time.Tick(interval):
			intervalTicked = true
		case <-manualReminder:
		}

		if task == nil {
			beeep.Notify("[active-task] No task", "No current task, assign a task", "a")
		} else {
			beeep.Notify("[active-task] #"+strconv.Itoa(taskId)+" "+task.Subject,
				"Started "+humanize.Time(startTime), "a")
		}

		if intervalTicked && i+1 != len(intervals) {
			i++
		}
	}

}
