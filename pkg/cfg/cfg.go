package cfg

import (
	"context"
	"fmt"
	"os"

	"github.com/PennState/sso2aws/pkg/envcfg"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
)

type Cfg struct {
	Profile   string    `config:"profile"`
	SSOConfig SSOConfig `config:"sso_config" yaml:"sso_config"`
}

type SSOConfig struct {
	SSOStartURL  string `config:"aws_sso_start_url" ini:"aws_sso_start_url" yaml:"aws_sso_start_url"`
	SSORegion    string `config:"aws_sso_region" ini:"aws_sso_region"`
	SSOAccountID int    `config:"aws_sso_account_id" ini:"aws_sso_account_id"`
	SSORoleName  string `config:"aws_sso_role_name" ini:"aws_sso_role_name"`
	Region       string `config:"aws_region" ini:"aws_region"`
	Output       string `config:"aws_output" ini:"aws_output"`
}

const (
	DefaultProfile = "sso2aws"
	DefaultRole    = "AdministratorAccess"
	DefaultRegion  = "us-east-1"
	DefaultOutput  = "json"

	EnvPrefix = "SSO2AWS"
)

func NewCfg() *Cfg {
	return &Cfg{
		SSOConfig: SSOConfig{
			SSORoleName: DefaultRole,
			SSORegion:   DefaultRegion,
			Region:      DefaultRegion,
			Output:      DefaultOutput,
		},
		Profile: DefaultProfile,
	}
}

func (cfg *Cfg) Load() error {
	loader := confita.NewLoader(
		file.NewOptionalBackend(defaultConfigFile()),
		envcfg.NewBackend(EnvPrefix),
	)

	return loader.Load(context.TODO(), cfg)
}

func defaultConfigDir() string {
	d, err := os.UserConfigDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not determine UserConfigDir: %s\n", err)
		os.Exit(1)
	}
	return d
}

func defaultConfigFile() string {
	return fmt.Sprintf("%s/sso2aws.yaml", defaultConfigDir())
}

func (cfg *Cfg) Validate() error {
	// check required keys

	return nil
}
