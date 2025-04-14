package main

import (
	"github.com/alecthomas/kong"
	"github.com/maikelh/navis/cmd"
	"github.com/maikelh/navis/internal"
)

type CLI struct {
	internal.Globals

	Apply   cmd.ApplyCmd   `cmd:"" help:"apply settings from config file"`
	Version cmd.VersionCmd `cmd:"" help:"Show the version information"`
}

func main() {
	cli := CLI{
		Globals: internal.Globals{},
	}

	ctx := kong.Parse(&cli,
		kong.Name("navis"),
		kong.Description("Tool to remotely deploy docker compose files"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
	)
	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}
