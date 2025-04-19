package internal

import (
	"fmt"
	"github.com/alecthomas/kong"
	"log/slog"
)

type Globals struct {
	Debug debugFlag `short:"D" help:"Enable debug mode"`
}

// A flag with a hook that, if triggered, will set the debug loggers output to stdout.
type debugFlag bool

func (d debugFlag) AfterApply() error {
	if d {
		SetupLogger(slog.LevelDebug)
	}
	return nil
}

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}
