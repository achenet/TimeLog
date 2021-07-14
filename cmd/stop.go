package cmd

import (
	"fmt"
	"github.com/achenet/TimeLog/data_store"
	"github.com/spf13/cobra"
	"time"
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
	tl.TaskList
}
