package octane

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigInitDefaults(t *testing.T) {
	cfg := Config{}

	err := cfg.InitDefaults()

	assert.Nil(t, err)

	phpPath, err := exec.LookPath("php")

	if err != nil {
		t.Error("error", err)
	}

	workingDirectory, err := os.Getwd()

	if err != nil {
		t.Error("error", err)
	}

	assert.Equal(t, workingDirectory, cfg.AppBasePath)
	assert.Equal(t, "production", cfg.Environment)
	assert.Equal(t, phpPath, cfg.PHPBinary)
}
