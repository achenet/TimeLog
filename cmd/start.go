// This command starts counting time on a task.
package cmd

import (
	"fmt"
	"github.com/achenet/TimeLog/data_store"
	"github.com/spf13/cobra"
	"time"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start logging time for a given task.",
	Run:   StartCmd,
}

func StartCmd(cmd *cobra.Command, args []string) {
	// Get args
	newTaskName := args[0]
	tl.DataStore.TaskList[newTaskName] = data_store.Task{
		Name:      newTaskName,
		TotalTime: time.Duration(0),
	}
	fmt.Println("Added task:", newTaskName)
}
