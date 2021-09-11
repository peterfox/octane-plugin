package octane

import (
	"github.com/spiral/errors"
	"github.com/spiral/roadrunner/v2/plugins/config"
)

// PluginName is the name for the plugin as found inside the container
const PluginName = "octane"

type Plugin struct {
	cfg    Config
}

// Init initiates the plugin with any injected services implementing endure.Container, returning an error if the
// plugin fails to start, if the error is of type errors.Disabled then the plugin will not be active
func (p *Plugin) Init(cfg config.Configurer) error {
	const op = errors.Op("octane_plugin_init")

	// if the config does not have a section matching PluginName section
	// then return an error
	if !cfg.Has(PluginName) {
		return errors.E(op, errors.Disabled)
	}

	// read in the section of the config by the plugin name
	// if it cannot be read then return an error
	err := cfg.UnmarshalKey(PluginName, &p.cfg)
	if err != nil {
		return errors.E(op, errors.Disabled, err)
	}

	err = p.cfg.InitDefaults()
	if err != nil {
		return errors.E(op, errors.Disabled, err)
	}

	// if the config does not specify it is enabled then discontinue
	if !p.cfg.Enabled {
		return errors.E(op, errors.Disabled)
	}

	err = cfg.Overwrite(map[string]interface{}{
		"server.command": p.cfg.PHPBinary + " ./vendor/bin/roadrunner-worker",
		"server.env": map[string]interface{}{
			"APP_ENVIRONMENT": p.cfg.Environment,
			"LARAVEL_OCTANE": 1,
			"APP_BASE_PATH": p.cfg.AppBasePath,
		},
	})

	if err != nil {
		return errors.E(op, errors.Disabled, err)
	}

	return nil
}

// Name returns endure.Named interface implementation
func (p *Plugin) Name() string {
	return PluginName
}
