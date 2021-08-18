package cmd

import (
	"github.com/Versent/saml2aws/pkg/shell"
	"github.com/apex/log"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:          "login",
	Short:        "login",
	Long:         "Login to a AWS SSO profile",
	RunE:         runLogin,
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func runLogin(cmd *cobra.Command, args []string) error {
	filename, err := updateAwsConfig()
	if err != nil {
		return err
	}

	log.Infof("Wrote AWS config profile '%s' to '%s'", c.Profile, filename)

	log.Infof("Running `aws sso login --profile=%s`", c.Profile)

	return shell.ExecShellCmd([]string{"aws", "sso", "login", "--profile=" + c.Profile}, []string{})
}
