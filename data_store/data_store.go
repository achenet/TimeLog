// The data_store package store data on the different tasks.
// Store information in a FOLDER, with each file being a task?
package data_store

import (
	"encoding/json"
	"os"
	"time"
	"fmt"
)

type DataStore struct {
	TaskList
	Storage string
}

func NewDataStore(taskList TaskList, storage string) *DataStore {
	return &DataStore{
		taskList,
		storage,
	}
}

type TaskList map[string]Task

type Task struct {
	Name         string // Should always be the same as the key for this task in the map.
	TotalTime    time.Duration
	CurrentStart time.Time
	Finished     bool
}

func (tl TaskList) CheckThatAllNamesAreCorrect() bool {
	for taskName, task := range tl {
		if task.Name != taskName {
			return false
		}
	}
	return true
}


func WriteToFile(file *os.File, data interface{}) error {
	// Marshal to json
	marshaledData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = file.Write(marshaledData)
	if err != nil {
		return err
	}
	fmt.Println("Wrote to file", file)
	return nil
}
