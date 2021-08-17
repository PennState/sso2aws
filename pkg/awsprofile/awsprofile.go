package awsprofile

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/PennState/sso2aws/pkg/cfg"
	"github.com/apex/log"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/go-ini/ini"
)

type AWSProfile struct {
	f        *ini.File
	Filename string
}

func DefaultSharedConfigFile() string {
	return config.DefaultSharedConfigFilename()
}

func Load() (AWSProfile, error) {
	ac := AWSProfile{Filename: config.DefaultSharedConfigFilename()}

	// if config file does not exist, create it
	if _, err := os.Stat(ac.Filename); err != nil {
		if os.IsNotExist(err) {
			log.Infof("File does not exist, creating it: %s", ac.Filename)

			dirPath := filepath.Dir(ac.Filename)

			err := os.Mkdir(dirPath, 0700)
			if err != nil {
				if !os.IsExist(err) {
					return ac, fmt.Errorf("unable to create '%s' directory: %s", dirPath, err)
				}
			}

			_, err = os.Create(ac.Filename)
			if err != nil {
				return ac, fmt.Errorf("unable to create configuration: %s", err)
			}
		}
	}

	f, err := ini.Load(ac.Filename)
	if err != nil {
		return ac, err
	}

	ac.f = f

	return ac, nil
}

func (ac AWSProfile) SetSSOProfile(name string, profile *cfg.SSOConfig) error {
	section, err := ac.f.NewSection("profile " + name)
	if err != nil {
		return err
	}

	return section.ReflectFrom(profile)
}

func (ac AWSProfile) Save() error {
	return ac.f.SaveTo(ac.Filename)
}
