package main

import (
	"github.com/alecthomas/kong"
	"github.com/maikelh/navis/cmd"
	"github.com/maikelh/navis/internal"
)

func main() {
	cli := cmd.CLI{
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
