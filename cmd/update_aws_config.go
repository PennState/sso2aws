package cmd

import (
	"fmt"

	"github.com/PennState/sso2aws/pkg/awsprofile"
	"github.com/apex/log"
	"github.com/spf13/cobra"
)

var updateAwsConfigCmd = &cobra.Command{
	Use:          "update-aws-config",
	Short:        "write aws profile to shared config",
	Long:         "write aws profile to shared config",
	RunE:         runUpdateAwsConfig,
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(updateAwsConfigCmd)
}

func runUpdateAwsConfig(cmd *cobra.Command, args []string) error {
	filename, err := updateAwsConfig()
	if err != nil {
		return err
	}

	log.Infof("Wrote AWS config profile '%s' to '%s'", c.Profile, filename)

	return nil
}

func updateAwsConfig() (string, error) {
	ac, err := awsprofile.Load()
	if err != nil {
		return "", fmt.Errorf("error loading AWS config: %s", err)
	}

	err = ac.SetSSOProfile(c.Profile, &c.SSOConfig)
	if err != nil {
		return "", fmt.Errorf("could not set SSO profile: %s", err)
	}

	err = ac.Save()
	if err != nil {
		return ac.Filename, fmt.Errorf("error saving AWS config profile '%s' to file '%s': %s", c.Profile, ac.Filename, err)
	}

	return ac.Filename, nil
}
