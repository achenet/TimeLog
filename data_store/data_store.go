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
			Name: taskName,
		}
	}

	// Set start time
	ds[taskName].CurrentStart = time.Now()
	ds[taskName].InProgress = true
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
		fmt.Println(name+", Total Time:", task.TotalTime.Minutes(), "minutes, in progress:", task.InProgress)
	}
}

func (ds DataStore) GetTaskNames() []string {
	taskNames := make([]string, 0, len(ds))
	for nameOfTask := range ds {
		taskNames = append(taskNames, nameOfTask)
	}
	return taskNames
}
