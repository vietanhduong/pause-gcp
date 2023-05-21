package main

import (
	"github.com/spf13/cobra"
	"github.com/vietanhduong/pause-gcp/cmd/pause"
	"github.com/vietanhduong/pause-gcp/cmd/version"
	"os"
)

func newRootCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "pause-gcp",
	}
	cmd.AddCommand(version.NewCommand())
	cmd.AddCommand(pause.NewCommand())
	return cmd
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
