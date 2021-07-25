package data_store

import (
	"fmt"
	"time"
)

type DataStore map[string]*Task

type Task struct {
	Name         string // Should always be the same as the key for this task in the map.
	TotalTime    time.Duration
	CurrentStart time.Time
	InProgress  bool
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
		return
	}

	// Make sure task is in progress (can't stop an already stopped task).
	if !ds[taskName].InProgress {
		fmt.Println("Task is not currently in progress and so cannot be stopped.")
		return
	}

	elapsedTime := time.Since(ds[taskName].CurrentStart)
	ds[taskName].TotalTime = ds[taskName].TotalTime + elapsedTime
}

// Show information contained in the datastore in a user-friendly format.
func (ds DataStore) ShowInfo() {
	for name, task := range ds {
		fmt.Println(name + ", Total Time:", task.TotalTime.Minutes(), "minutes, in progress:", task.InProgress)
	}
}
