// This command starts counting time on a task.
package cmd

import (
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start logging time for a given task.",
	Run:   StartCmd,
}

func StartCmd(cmd *cobra.Command, args []string) {
	// Get task
	taskName := getTaskName(args)

	// Get file
	ds := readAndUnmarshalStorageFile()
	ds.StartTask(taskName)

	// Write to file
	marshalAndWriteToStorageFile(ds)
}
