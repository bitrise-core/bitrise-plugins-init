package cli

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/bitrise-core/bitrise-init/scanner"
	envmanModels "github.com/bitrise-io/envman/models"
	"github.com/bitrise-io/go-utils/fileutil"
	"github.com/bitrise-io/go-utils/pathutil"
	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v1"
)

var configCmd = cli.Command{
	Name:  "config",
	Usage: "Scans your project and generates bitrise config and secrets",
	Action: func(c *cli.Context) {
		if err := config(c); err != nil {
			log.Fatal(err)
		}
	},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Usage: "bitrise config file path",
			Value: "bitrise.yml",
		},
		cli.StringFlag{
			Name:  "secrets",
			Usage: "bitrise secrets file path",
			Value: ".bitrise.secrets.yml",
		},
	},
}

func config(c *cli.Context) error {
	// validate inputs
	configPth := c.String("config")
	secretsPth := c.String("secrets")

	if configPth == "" {
		return fmt.Errorf("config path not specified")
	}

	if exist, err := pathutil.IsPathExists(configPth); err != nil {
		return err
	} else if exist {
		return fmt.Errorf("config path (%s) already exist", configPth)
	}

	if secretsPth == "" {
		return fmt.Errorf("secrets path not specified")
	}

	if exist, err := pathutil.IsPathExists(secretsPth); err != nil {
		return err
	} else if exist {
		return fmt.Errorf("secrets path (%s) already exist", secretsPth)
	}

	// run scanner
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory, error: %s", err)
	}

	scanResult, err := scanner.Config(currentDir)
	if err != nil {
		return err
	}

	if len(scanResult.OptionsMap) == 0 {
		return fmt.Errorf("No known platform type detected")
	}

	bitriseConfig, err := scanner.AskForConfig(scanResult)
	if err != nil {
		return err
	}

	// write outputs
	configDir := filepath.Dir(configPth)
	if exist, err := pathutil.IsDirExists(configDir); err != nil {
		return err
	} else if !exist {
		if err := os.MkdirAll(configDir, 0700); err != nil {
			return fmt.Errorf("failed to create config directory (%s), error: %s", configDir, err)
		}
	}

	secretsDir := filepath.Dir(secretsPth)
	if exist, err := pathutil.IsDirExists(secretsDir); err != nil {
		return err
	} else if !exist {
		if err := os.MkdirAll(secretsDir, 0700); err != nil {
			return fmt.Errorf("failed to create secrets directory (%s), error: %s", secretsDir, err)
		}
	}

	configBytes, err := yaml.Marshal(bitriseConfig)
	if err != nil {
		return fmt.Errorf("failed to marshal bitrise config, error: %s", err)
	}

	if err := fileutil.WriteBytesToFile(configPth, configBytes); err != nil {
		return fmt.Errorf("failed to write bitrise config, error: %s", err)
	}

	log.Infof("bitrise config generated at: %s", configPth)

	secrets := envmanModels.EnvsYMLModel{}

	secretsBytes, err := yaml.Marshal(secrets)
	if err != nil {
		return fmt.Errorf("failed to marshal bitrise secrets, error: %s", err)
	}

	if err := fileutil.WriteBytesToFile(secretsPth, secretsBytes); err != nil {
		return fmt.Errorf("failed to write bitrise secrets, error: %s", err)
	}

	log.Infof("bitrise secrets generated at: %s", secretsPth)

	return nil
}
