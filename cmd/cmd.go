package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/achenet/TimeLog/data_store"
	"os"
)

const (
	DirectoryName   = "/.timelog"
	StorageFileName = "datastore.json"
)

var (
	StorageFilePath  string
	StorageDirectory string
)

func SetStorageFilePath() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err.Error())
		os.Exit(1)
	}

	StorageDirectory = homeDir + DirectoryName
	StorageFilePath = StorageDirectory + StorageFileName
}

func readAndUnmarshalStorageFile() data_store.DataStore {

	// Read file
	timeLogData, err := os.ReadFile(StorageFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading file: "+err.Error())
		os.Exit(1)
	}

	// Unmarshal
	var tl data_store.DataStore
	err = json.Unmarshal(timeLogData, &tl)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error unmarshaling data: "+err.Error())
		os.Exit(1)
	}
	return tl
}

func marshalAndWriteToStorageFile(data interface{}) {
	// Use os.create to truncate file
	file, err := os.Create(StorageFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating or truncating file: "+err.Error())
		os.Exit(1)

	}

	// Marshal data
	marshaledData, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error marshaling data: "+err.Error())
		os.Exit(1)
	}

	// Write to file
	_, err = file.Write(marshaledData)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error writing to file: "+err.Error())
		os.Exit(1)
	}
}

func getTaskName(args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "No task name specified.")
		os.Exit(1)
	}
	return args[0]
}
