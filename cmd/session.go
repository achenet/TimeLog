package cmd

import (
	"github.com/spf13/cobra"
)

var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Show task sessions.",
	Run:   SessionCmd,
}

func SessionCmd(cmd *cobra.Command, args []string) {
	// Get args
	taskName := getTaskName(args)

	// Get file
	ds := readAndUnmarshalStorageFile()

	// Show sessions
	ds.ShowTaskSession(taskName)

}
