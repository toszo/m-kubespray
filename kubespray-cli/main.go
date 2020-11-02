package main

import (
	"github.com/rs/zerolog"
	"github.com/toszo/kubespray-cli/cmd"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	cmd.Execute()
}
