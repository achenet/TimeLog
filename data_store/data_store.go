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
}

func (ds DataStore) CheckThatAllNamesAreCorrect() bool {
	for taskName, task := range ds {
		if task.Name != taskName {
			return false
		}
	}
	return true
}

func (ds DataStore) StartTask(taskName string) {
	// Check if task exists, create if not.
	if _, ok := ds[taskName]; !ok {
		ds[taskName] = &Task{
			Name: taskName,
		}
	}

	// Set start time
	ds[taskName].CurrentStart = time.Now()
}

func (ds DataStore) StopTask(taskName string) {
	// Make sure the task exists
	if _, ok := ds[taskName]; !ok {
		fmt.Println("No task with that name was found.")
		return
	}

	elapsedTime := time.Since(ds[taskName].CurrentStart)
	ds[taskName].TotalTime = ds[taskName].TotalTime + elapsedTime
}
