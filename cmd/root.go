package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/achenet/TimeLog/data_store"
	"github.com/spf13/cobra"
	"os"
)

const storageFilePath = "~/.timelog/datastore.json"

var rootCmd = &cobra.Command{
	Use:   "timelog",
	Short: "A simple program to log time spent on various tasks.",
	Run:   RootCmd,
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
}

// Usage timelog start taskName: start counting on the task. Create it if it doesn't exist.
// Usage timelog stop taskName: stop counting on the task. Return error if it doesn't exist.
// Usage timeog: print all tasks, with total time for each.
func RootCmd(cmd *cobra.Command, args []string) {
	tl := readAndUnmarshalStorageFile()
	fmt.Println(tl)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func readAndUnmarshalStorageFile() *data_store.DataStore {

	// Read file
	timeLogData, err := os.ReadFile(storageFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// Unmarshal
	var tl data_store.DataStore
	err := json.Unmarshal(timeLogData, &tl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
