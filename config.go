package octane

import (
	"os"
	"os/exec"
)

// Config configures metrics service.
type Config struct {
	// the Enabled property makes the plugin exit prematurely
	Enabled     bool `mapstructure:"enable"`
	// the PHPBinary is the php executable path
	PHPBinary   string `mapstructure:"php_binary"`
	// the AppBasePath is the path of the existing php project
	AppBasePath string `mapstructure:"app_path"`
	// the Environment is the Laravel environment to run the working in
	Environment string
}

// InitDefaults for the octane config
func (cfg *Config) InitDefaults() error {
	if cfg.Enabled != true {
		cfg.Enabled = false
	}

	if cfg.PHPBinary == "" {
		phpPath, err := exec.LookPath("php")
		if err != nil {
			return err
		}

		cfg.PHPBinary = phpPath
	}

	if cfg.AppBasePath == "" {
		workingDirectory, err := os.Getwd()

		if err != nil {
			return err
		}

		cfg.AppBasePath = workingDirectory
	}

	if cfg.Environment == "" {
		cfg.Environment = "production"
	}

	return nil
}
