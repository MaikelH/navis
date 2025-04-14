package internal

import (
	"fmt"
	"github.com/alecthomas/kong"
)

type Globals struct {
	Debug bool `short:"D" help:"Enable debug mode"`
}

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}
