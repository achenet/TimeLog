// The data_store package store data on the different tasks.
// Store information in a FOLDER, with each file being a task?
package data_store

import (
	"time"
)

type DataStore struct {
	TaskList
	Storage string
}


type TaskList map[string]Task

type Task struct {
	Name         string // Should always be the same as the key for this task in the map.
	TotalTime    time.Duration
	CurrentStart time.Time
	Finished     bool
}

func NewDataStore(taskList TaskList, storage string) *DataStore {
	return &DataStore{
		taskList,
		storage,
	}
}

func (tl TaskList) CheckThatAllNamesAreCorrect() bool {
	for taskName, task := range tl {
		if task.Name != taskName {
			return false
		}
	}
	return true
}

