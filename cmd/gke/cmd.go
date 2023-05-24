package gke

import (
	"github.com/spf13/cobra"
	"github.com/vietanhduong/pause-gcp/cmd/gke/pause"
	"github.com/vietanhduong/pause-gcp/cmd/gke/refresh"
	"github.com/vietanhduong/pause-gcp/cmd/gke/unpause"
)

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "gke",
		Short: "Manage GKE resource",
		Long:  `The gke command lets you can mange your gke cluster like pause/unpause or refresh.`,
	}

	cmd.AddCommand(pause.NewCommand())
	cmd.AddCommand(unpause.NewCommand())
	cmd.AddCommand(refresh.NewCommand())
	return cmd
}
