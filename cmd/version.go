package cmd

import (
	"fmt"
	"github.com/maikelh/navis/internal"
)

type VersionCmd struct {
}

func (cmd *VersionCmd) Run(globals *internal.Globals) error {
	fmt.Println("navis 0.0.1")
	return nil
}
