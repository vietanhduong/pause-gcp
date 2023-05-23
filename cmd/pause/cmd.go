package pause

import (
	"github.com/spf13/cobra"
	"github.com/vietanhduong/pause-gcp/cmd/pause/gke"
)

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "pause",
		Short: "Pause GCP resources",
		Long: `Pause GCP resources.
Currently, we support GKE, Virtual Machine and Cloud Sql.`,
	}

	cmd.AddCommand(gke.NewCommand())
	return cmd
}
