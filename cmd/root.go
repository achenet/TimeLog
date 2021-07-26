package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "timelog",
	Short: "A simple program to log time spent on various tasks.",
	Run:   RootCmd,
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(deleteCmd)
}

// Usage timelog start taskName: start counting on the task. Create it if it doesn't exist.
// Usage timelog stop taskName: stop counting on the task. Return error if it doesn't exist.
// Usage timeog: print all tasks, with total time for each.
func RootCmd(cmd *cobra.Command, args []string) {
	ds := readAndUnmarshalStorageFile()
	ds.ShowInfo()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
