package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/vietanhduong/pause-gcp/cmd/gke"
	"github.com/vietanhduong/pause-gcp/cmd/sql"
	"github.com/vietanhduong/pause-gcp/cmd/version"
	"github.com/vietanhduong/pause-gcp/cmd/vm"
)

func newRootCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "pause-gcp",
	}
	cmd.AddCommand(version.NewCommand())
	cmd.AddCommand(gke.NewCommand())
	cmd.AddCommand(vm.NewCommand())
	cmd.AddCommand(sql.NewCommand())
	return cmd
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
