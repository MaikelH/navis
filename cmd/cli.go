package cmd

import "github.com/maikelh/navis/internal"

type CLI struct {
	internal.Globals

	Apply   ApplyCmd   `cmd:"" help:"apply settings from config file"`
	Version VersionCmd `cmd:"" help:"Show the version information"`
}
