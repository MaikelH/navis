package cmd

import (
	"github.com/maikelh/navis/internal"
	"log/slog"
)

type ApplyCmd struct {
	ConfigFile string `help:"config file" default:"./navis.yaml" type:"path"`
}

func (cmd *ApplyCmd) Run(globals *internal.Globals) error {
	slog.Info("Applying config file", "file", cmd.ConfigFile)

	return nil
}
