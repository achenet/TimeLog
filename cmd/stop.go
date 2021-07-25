package cmd

import (
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop logging time for a given task.",
	Run:   StopCmd,
}

func StopCmd(cmd *cobra.Command, args []string) {
	// Get args
	taskName := args[0]

	// Get file
	tl := readAndUnmarshalStorageFile()

	// Stop task
	tl.StopTask(taskName)

	marshalAndWriteToStorageFile(tl)
}
