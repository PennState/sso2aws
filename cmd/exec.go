package cmd

import (
	"github.com/Versent/saml2aws/pkg/shell"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:          "exec [cmd] [args...]",
	Short:        "execute commands with the configured profile",
	Long:         "execute commands with the configured profile",
	RunE:         runExec,
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(execCmd)
}

func runExec(cmd *cobra.Command, args []string) error {
	_, err := updateAwsConfig()
	if err != nil {
		return err
	}

	shell.ExecShellCmd(args, []string{"AWS_PROFILE=" + c.Profile})

	return nil
}
