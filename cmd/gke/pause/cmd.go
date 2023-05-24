package pause

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
		Use:   "pause [CLUSTER_NAME]",
		Short: "Pause a GKE cluster",
		Long: `Pause a GKE cluster.
This command require '--location' and '--project' flags.`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 || len(args[0]) == 0 {
				return errors.Errorf("CLUSTER_NAME is required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := utils.RequiredFlags(cmd, "location", "project"); err != nil {
				return err
			}
			runCfg.clusterName = args[0]
			return run(runCfg)
		},
	}

	cmd.Flags().StringVarP(&runCfg.location, "location", "l", "asia-southeast1", "the cluster location")
	cmd.Flags().StringVarP(&runCfg.project, "project", "p", "", "the project where contain the cluster")
	cmd.Flags().StringVar(&runCfg.outputDir, "output-dir", "", "the output directory to write the cluster state. If no path is specified, this will skip the write-to-file process. The output state file has named by format `gke_<project>_<location>_<cluster_name>.json`")
	cmd.Flags().StringSliceVar(&runCfg.exceptPools, "except-pools", nil, "except node pools")
	return cmd
}
