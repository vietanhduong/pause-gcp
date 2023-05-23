package unpause

import (
	"github.com/spf13/cobra"
	"github.com/vietanhduong/pause-gcp/cmd/unpause/gke"
)

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "unpause",
		Short: "Unpause GCP resources",
		Long:  `Unpause GCP resources.`,
	}

	cmd.AddCommand(gke.NewCommand())
	return cmd
}
