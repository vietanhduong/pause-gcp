package refresh

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/vietanhduong/pause-gcp/cmd/utils"
)

func NewCommand() *cobra.Command {
	var (
		runCfg runConfig
	)

	var cmd = &cobra.Command{
		Use:   "refresh [CLUSTER_NAME]",
		Short: "Refresh a GKE cluster",
		Long: `Refresh a GKE cluster.
Refresh all worker nodes in all node pools of the input cluster. This command just works with node pool has type is spot or preemptible.`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.Errorf("CLUSTER_NAME is missing")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := utils.RequiredFlags(cmd, "location", "project"); err != nil {
				return err
			}
			runCfg.name = args[0]
			return run(runCfg)
		},
	}

	cmd.Flags().StringVarP(&runCfg.location, "location", "l", "asia-southeast1", "the cluster location")
	cmd.Flags().StringVarP(&runCfg.project, "project", "p", "", "the project where contain the cluster")
	cmd.Flags().BoolVar(&runCfg.recreate, "recreate", false, "keep the instance (node) name or delete and create with new name otherwise")

	return cmd
}
