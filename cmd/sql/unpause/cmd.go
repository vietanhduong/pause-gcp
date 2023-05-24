package unpause

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
		Use:   "unpause [INSTANCE_NAME]",
		Short: "Unpause a SQL instance",
		Long: `Unpause a SQL instance.
This command require '--project' flags.`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 || len(args[0]) == 0 {
				return errors.Errorf("INSTANCE_NAME is required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := utils.RequiredFlags(cmd, "project"); err != nil {
				return err
			}
			runCfg.name = args[0]
			return run(runCfg)
		},
	}

	cmd.Flags().StringVarP(&runCfg.project, "project", "p", "", "the project where contain the SQL instance")
	return cmd
}
