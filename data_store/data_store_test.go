package data_store

import (
	"fmt"
	"os"
	"testing"
)

const testStorage = "./testStorage"

var testDataStore *DataStore

func TestMain(m *testing.M) {
	fmt.Println("Testing data store module")
	var taskList TaskList
	var storage = testStorage
	testDataStore = NewDataStore(taskList, storage)
	m.Run()
}

func TestWriteToFile(t *testing.T) {
	fmt.Println("Testing WriteToFile")
	testToWrite := TaskList{
		"snoop": {
			Name: "snoop",
		},
		"jack": {
			Name: "jack",
		},
	}
	testDataStore.TaskList = testToWrite
	file, err := os.Create(testStorage)
	if err != nil {
		t.Error(err.Error())
	}
	testDataStore.WriteToFile(file, testToWrite)
}

func TestReadFromFile(t *testing.T) {
	// Would first have to write test file
}
