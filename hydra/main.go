//go:generate swagger generate spec

package main

import (
	"github.com/justpark/auth/hydra/cmd"
	"github.com/ory/x/profilex"
)

func main() {
	defer profilex.Profile().Stop()

	cmd.Execute()
}
