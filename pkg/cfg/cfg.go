package cfg

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
	SSOStartURL  string `config:"sso_start_url" ini:"sso_start_url"`
	SSORegion    string `config:"sso_region" ini:"sso_region"`
	SSOAccountID int    `config:"sso_account_id" ini:"sso_account_id"`
	SSORoleName  string `config:"sso_role_name" ini:"sso_role_name"`
	Region       string `config:"region" ini:"region"`
	Output       string `config:"output" ini:"output"`
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

func (cfg *Cfg) Print() {
	j, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(j))
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
