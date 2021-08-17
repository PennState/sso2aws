package cmd

import (
	"fmt"
	"os"

	"git.psu.edu/swe-golang/buildversion"
	"github.com/PennState/sso2aws/pkg/cfg"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:          os.Args[0],
	Short:        "manage AWS SSO CLI environments",
	Long:         "A tool to make operating under multiple AWS SSO accounts and roles easier",
	SilenceUsage: true,
	Version:      fmt.Sprintf("%+v", buildversion.Get()),
}

var c *cfg.Cfg

func init() {
	log.SetHandler(cli.Default)

	debug := flag.BoolP("debug", "D", false, "debug output")
	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	c = cfg.NewCfg()
	err := c.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}
}

// Execute executes this command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
