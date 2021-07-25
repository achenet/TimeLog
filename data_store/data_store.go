package data_store

import (
	"time"
)

type DataStore map[string]*Task

type Task struct {
	Name         string // Should always be the same as the key for this task in the map.
	TotalTime    time.Duration
	CurrentStart time.Time
}

func (tl DataStore) CheckThatAllNamesAreCorrect() bool {
	for taskName, task := range tl {
		if task.Name != taskName {
			return false
		}
	}
	return true
}

func (tl DataStore) StopTask(taskName string) {
	elapsedTime := time.Since(tl[taskName].CurrentStart)
	tl[taskName].TotalTime = tl[taskName].TotalTime + elapsedTime
}