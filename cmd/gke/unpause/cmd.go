package unpause

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	apis "github.com/vietanhduong/pause-gcp/apis/v1"
	"github.com/vietanhduong/pause-gcp/pkg/utils/protoutil"
	"os"
)

func NewCommand() *cobra.Command {
	var (
		runCfg runConfig
		rm     bool
	)

	var cmd = &cobra.Command{
		Use:   "unpause [STATE_FILE]",
		Short: "Unpause a GKE cluster",
		Long: `Unpause a GKE cluster.
This command requires a GKE state file which is created when you pause the cluster.`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.Errorf("STATE_FILE is missing")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			b, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}
			var cluster apis.Cluster
			if err = protoutil.Unmarshal(b, &cluster); err != nil {
				return err
			}
			runCfg.cluster = &cluster
			if err = run(runCfg); err != nil {
				return err
			}
			if rm {
				return os.Remove(args[0])
			}
			return nil
		},
	}

	cmd.Flags().BoolVar(&rm, "rm", false, "Remove the cluster state after complete")

	return cmd
}
