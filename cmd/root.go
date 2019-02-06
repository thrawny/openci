package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thrawny/openci/pkg/open"
)

var rootCmd = &cobra.Command{
	Use:   "openci",
	Short: "Run the ci build url of a repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		return open.Run()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
