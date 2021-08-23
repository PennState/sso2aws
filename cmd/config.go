package cmd

import (
	"github.com/spf13/cobra"
)

var getConfigCmd = &cobra.Command{
	Use:          "get-config",
	Short:        "show runtime config",
	Long:         "show runtime config",
	RunE:         runGetConfig,
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(getConfigCmd)
}

func runGetConfig(cmd *cobra.Command, args []string) error {
	c.Print()

	return nil
}
