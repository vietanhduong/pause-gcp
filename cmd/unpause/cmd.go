package unpause

import (
	"github.com/spf13/cobra"
	"github.com/vietanhduong/pause-gcp/cmd/utils"
	"github.com/vietanhduong/pause-gcp/pkg/utils/exec"
)

func NewCommand() *cobra.Command {
	var (
		runCfg runConfig
	)

	var cmd = &cobra.Command{
		Use:   "unpause",
		Short: "Unpause GCP resources",
		Long: `Unpause GCP resources.
This command require a config file to detect the backup state folder. If there is no backup state found, no resource is unpause.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := utils.RequiredFlags(cmd, "config"); err != nil {
				return err
			}
			if err := testGcloud(); err != nil {
				return err
			}
			return run(runCfg)
		},
	}

	cmd.Flags().StringVarP(&runCfg.configFile, "config", "c", "", "Pause GCP's config file. The input file must be one of `yaml` or `json` format.")
	return cmd
}

func testGcloud() error {
	_, err := exec.Run(exec.Command("which", "gcloud"))
	return err
}
