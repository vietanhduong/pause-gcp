package vm

import (
	"github.com/spf13/cobra"
	"github.com/vietanhduong/pause-gcp/cmd/vm/pause"
	"github.com/vietanhduong/pause-gcp/cmd/vm/unpause"
)

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "vm",
		Short: "Manage Compute instance resources",
		Long:  `The vm command lets you mange Compute Engine instances like pause/unpause.`,
	}

	cmd.AddCommand(pause.NewCommand())
	cmd.AddCommand(unpause.NewCommand())
	return cmd
}
