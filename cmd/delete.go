package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the task.",
	Run:   DeleteCmd,
}

func DeleteCmd(cmd *cobra.Command, args []string) {
	// Get args
	taskName := args[0]

	// Get file
	tl := readAndUnmarshalStorageFile()

	// Delete task
	tl.DeleteTask(taskName)

	marshalAndWriteToStorageFile(tl)
}
