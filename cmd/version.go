package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:          "version",
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", rootCmd.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
