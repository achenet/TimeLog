package data_store

import (
	"encoding/json"
	"fmt"
	"os"
)

func (ds *DataStore) WriteToFile(file *os.File, data interface{}) error {
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

func (ds *DataStore) ReadFromFile(filename string) error {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return nil
	// Unmarshal
	var taskList TaskList
	err = json.Unmarshal(bytes, taskList)
	if err != nil {
		return err
	}
	ds.TaskList = taskList
	fmt.Println(ds.TaskList)
	return nil
}
