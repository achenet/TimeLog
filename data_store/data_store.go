package data_store

import (
	"fmt"
	"github.com/achenet/TimeLog/auto_suggest"
	"time"
)

type DataStore map[string]*Task

type Task struct {
	Name         string // Should always be the same as the key for this task in the map.
	TotalTime    time.Duration
	CurrentStart time.Time
	InProgress   bool
	Sessions     []*Session
}

type Session struct {
	Start    time.Time
	End      time.Time
	Duration time.Duration
}

func (ds DataStore) CheckThatAllNamesAreCorrect() bool {
	for taskName, task := range ds {
		if task.Name != taskName {
			return false
		}
	}
	return true
}

// Start counting time on a task.
func (ds DataStore) StartTask(taskName string) {
	// Check if task exists, create if not.
	if _, ok := ds[taskName]; !ok {
		ds[taskName] = &Task{
			Name:     taskName,
			Sessions: []*Session{},
		}
	}

	// Set start time
	startTime := time.Now()
	ds[taskName].CurrentStart = startTime
	ds[taskName].InProgress = true

	// Create new session
	ds[taskName].Sessions = append(ds[taskName].Sessions,
		&Session{
			Start: startTime,
		})
}

// Tell the time log software to stop counting time on an existing task
func (ds DataStore) StopTask(taskName string) {
	// Make sure the task exists
	if _, ok := ds[taskName]; !ok {
		fmt.Println("No task with that name was found.")
		// Find possible similar tasks using auto_suggest
		possibleTasks := auto_suggest.AutoSuggest(taskName, ds.GetTaskNames())
		fmt.Println("Possible tasks include:", possibleTasks)
		return
	}

	// Make sure task is in progress (can't stop an already stopped task).
	if !ds[taskName].InProgress {
		fmt.Println("Task is not currently in progress and so cannot be stopped.")
		return
	}

	elapsedTime := time.Since(ds[taskName].CurrentStart)
	ds[taskName].TotalTime += elapsedTime
	ds[taskName].InProgress = false

	// Close session
	ds[taskName].Sessions[len(ds[taskName].Sessions)-1].End = time.Now()
	ds[taskName].Sessions[len(ds[taskName].Sessions)-1].Duration = elapsedTime
}

// Delete a task if it exists
func (ds DataStore) DeleteTask(taskName string) {
	// Make sure the task exists
	if _, ok := ds[taskName]; !ok {
		fmt.Println("No task with that name was found.")
		// Find possible similar tasks using auto_suggest
		possibleTasks := auto_suggest.AutoSuggest(taskName, ds.GetTaskNames())
		fmt.Println("Possible tasks include:", possibleTasks)
		return
	}

	delete(ds, taskName)
}

// Show information contained in the datastore in a user-friendly format.
func (ds DataStore) ShowInfo() {
	for name, task := range ds {
		taskStr := fmt.Sprintf(name+", Total Time: %v minutes", task.TotalTime.Minutes())
		if task.InProgress {
			taskStr += ", in progress, current session started at:" + fmt.Sprintf(" %v", task.CurrentStart)
		}
		fmt.Println(taskStr)
	}
}

func (ds DataStore) GetTaskNames() []string {
	taskNames := make([]string, 0, len(ds))
	for nameOfTask := range ds {
		taskNames = append(taskNames, nameOfTask)
	}
	return taskNames
}

func (ds DataStore) ShowTaskSession(taskName string) {
	// Make sure the task exists
	if _, ok := ds[taskName]; !ok {
		fmt.Println("No task with that name was found.")
		// Find possible similar tasks using auto_suggest
		possibleTasks := auto_suggest.AutoSuggest(taskName, ds.GetTaskNames())
		fmt.Println("Possible tasks include:", possibleTasks)
		return
	}

	for _, session := range ds[taskName].Sessions {
		if session.Duration == 0 {
			fmt.Println("Start:", session.Start)
            return
		}
		fmt.Println("Start:", session.Start, "End:", session.End, "Duration:", session.Duration)
	}
}
