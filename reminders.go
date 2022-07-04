package activetask

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gen2brain/beeep"
)

const notWorkingTaskId = -1

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

func LogTime(startTime time.Time, task string) {
	if task != "" {
		beeep.Notify(fmt.Sprintf("[active-task] %s", task),
			"Logged "+humanize.RelTime(startTime, time.Now(), "", ""), "a")
	}
}

func fuzzDuration(duration time.Duration, fuzz float64) time.Duration {
	if rand.Intn(1) == 1 {
		fuzz = fuzz * -1
	}
	return duration + time.Duration(rand.Float64()*fuzz)
}

func parseIdFromTask(task string) int {
	// Any number which starts within 5 characters of the task message is the id
	r := regexp.MustCompile(".{0,5}([0-9]+)")
	submatches := r.FindStringSubmatch(task)
	if submatches[1] == "" {
		return notWorkingTaskId
	}
	id, err := strconv.Atoi(submatches[1])
	if err != nil {
		// Shouldn't be possible
		log.Printf("Error parsing ID from task message %s: %s", task, err.Error())
		return notWorkingTaskId
	}
	return id
}

// Start issuing reminders on a graduated interval, indicating the time when the task started
// If the task is null, remind the user that they have no current task
func IssueRemindersAndLogTime(startTime time.Time, task string, manualReminder chan bool, cancel chan bool) {
	var i, taskId int

	var intervals []time.Duration

	if task == "" {
		taskId = notWorkingTaskId
		intervals = graduatedNotificationIntervalsNoTask
	} else {
		taskId = parseIdFromTask(task)
		intervals = graduatedNotificationIntervals
	}
	i = 0
	// Initial ticker
	ticker := time.NewTicker(fuzzDuration(intervals[i], 0.15))
	for {

		intervalTicked := false

		select {
		case <-cancel:
			log.Printf("Aborting reminders for %d", taskId)
			LogTime(startTime, task)
			return
		case <-ticker.C:
			intervalTicked = true
		case <-manualReminder:
		}

		var err error
		if task == "" {
			err = beeep.Notify("[active-task] No task", "No current task, assign a task", "a")
		} else {
			err = beeep.Notify(fmt.Sprintf("[active-task] %s", task),
				"Started "+humanize.Time(startTime), "a")
		}
		if err != nil {
			log.Fatal(err)
		}

		if intervalTicked && i+1 != len(intervals) {
			i++
			ticker.Stop()
			ticker = time.NewTicker(fuzzDuration(intervals[i], 0.15))
		}
	}

}
