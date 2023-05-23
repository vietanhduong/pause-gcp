package gke

import (
	"github.com/spf13/cobra"
	"github.com/vietanhduong/pause-gcp/cmd/utils"
)

func NewCommand() *cobra.Command {
	var (
		runCfg runConfig
	)

	var cmd = &cobra.Command{
		Use:   "gke",
		Short: "Pause a GKE cluster",
		Long: `Pause a GKE cluster. 
This command require '--name', '--location' and '--project' flags.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := utils.RequiredFlags(cmd, "name", "location", "project"); err != nil {
				return err
			}
			return run(runCfg)
		},
	}

	cmd.Flags().StringVarP(&runCfg.clusterName, "name", "n", "", "The cluster name.")
	cmd.Flags().StringVarP(&runCfg.location, "location", "l", "asia-southeast1", "The cluster location. Could be a zone or a region.")
	cmd.Flags().StringVarP(&runCfg.project, "project", "p", "", "The project where contain the cluster.")
	cmd.Flags().StringVar(&runCfg.outputDir, "output-dir", ".", "The output directory to write the cluster state.")
	return cmd
}
