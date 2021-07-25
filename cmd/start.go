// This command starts counting time on a task.
package cmd

import (
	"github.com/spf13/cobra"
	"time"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start logging time for a given task.",
	Run:   StartCmd,
}

func StartCmd(cmd *cobra.Command, args []string) {
	taskName := args[0]

	tl := readAndUnmarshalStorageFile()
	tl[taskName].CurrentStart = time.Now()

	// Write to file
	marshalAndWriteToStorageFile(tl)
}
