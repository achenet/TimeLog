package main

import (
	"encoding/json"
	"fmt"
	"github.com/achenet/TimeLog/cmd"
	"github.com/achenet/TimeLog/data_store"
	"os"
)

func main() {

	// Set storage file path
	cmd.SetStorageFilePath()

	// Make storage directory if it isn't already created
	if err := os.MkdirAll(TaskListDirectory, 744); err != nil {
		fmt.Println("Error making storage directory: \n", err.Error())
		os.Exit(1)
	}

	// Create datastore file if it doesn't already exist
	fmt.Println("Opening file")
	_, err := os.Open(cmd.StorageFilePath)
	if err != nil {

		fmt.Println("Could not open file, due to", err.Error(), "\nOpting to create one instead.")

		file, err := os.Create(cmd.StorageFilePath)
		if err != nil {
			fmt.Println("Could not create storage file:\n" + err.Error())
			os.Exit(1)
		}

		// If it doesn't exist, write something empty in it.
		tl := make(data_store.DataStore)
		m, err := json.Marshal(tl)
		if err != nil {
			fmt.Println("Error creating blank file:\n" + err.Error())
			os.Exit(1)
		}

		_, err = file.Write(m)
		if err != nil {
			fmt.Println("Error writing empty datastore to new file:\n" + err.Error())
			os.Exit(1)
		}

		fmt.Println("Created and wrote to empty file")
	}

	cmd.Execute()
}
