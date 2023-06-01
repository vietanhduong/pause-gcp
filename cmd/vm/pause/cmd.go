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
		Use:   "pause [INSTANCE_NAME]",
		Short: "Pause a Virtual Machine",
		Long: `Pause a Virtual Machine. 
This command require '--zone' and '--project' flags.`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 || len(args[0]) == 0 {
				return errors.Errorf("INSTANCE_NAME is required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := utils.RequiredFlags(cmd, "zone", "project"); err != nil {
				return err
			}
			runCfg.name = args[0]
			return run(runCfg)
		},
	}

	cmd.Flags().StringVarP(&runCfg.zone, "zone", "z", "asia-southeast1-b", "the instance's zone")
	cmd.Flags().StringVarP(&runCfg.project, "project", "p", "", "the project where contain the instance")
	cmd.Flags().BoolVar(&runCfg.terminate, "terminate", false, "terminate the instance or suspend otherwise")
	return cmd
}
