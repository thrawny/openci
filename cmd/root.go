package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thrawny/openci/pkg/open"
	"os"
)

var Directory string

var rootCmd = &cobra.Command{
	Use:   "openci",
	Short: "Run the ci build url of a repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		return open.Run(Directory)
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&Directory, "directory", "d", "", "The directory open for. Defaults to the current working directory.")
}

func initConfig() {
	if Directory == "" {
		wd, err := os.Getwd()
		if err != nil {
			os.Exit(1)
		}
		Directory = wd
	}
}

func Execute() error {
	return rootCmd.Execute()
}
