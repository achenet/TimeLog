package main

import (
	"fmt"
	"github.com/achenet/TimeLog/cmd"
	"log"
	"os"
)

const (
	TaskListDirectory = "~/.taskList/"
)

var (
	tl       cmd.TimeLogger
	taskList data_store.TaskList
)

func init() {
	// Make storage directory if it isn't already created
	if err := os.MkdirAll(TaskListDirectory, 744); err != nil {
		log.Fatal("Error making storage directory:", err.Error())
	}

	// Read from storage directory if it has anything in it.
	storageDirectory, err := os.Open(TaskListDirectory)
	if err != nil {
		log.Fatal("Error opening storage directory:", err.Error())
	}
	storageDirectoryEntries, err := storageDirectory.ReadDir(-1)
	if err != nil {
		log.Fatal("Error reading storage directory", err.Error())
	}

	// If empty, finish init
	if len(storageDirectoryEntries) == 0 {
		return
	}

	for _, task := range storageDirectoryEntries {

	}
}

func main() {
	tl = NewTimeLogger(taskList)
	fmt.Println("TimeLog")
	tl.Execute()
}
