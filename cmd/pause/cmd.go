package pause

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
		Use:   "pause",
		Short: "Pause GCP resources",
		Long: `Pause GCP resources.
Currently, we support GKE, Virtual Machine and Cloud Sql.`,
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
	cmd.Flags().BoolVar(&runCfg.force, "force", false, "Pause GCP resources even if not in a schedule if this flag is presented.")
	cmd.Flags().BoolVar(&runCfg.dryRun, "dry-run", false, "Simulate a pause")
	return cmd
}

func testGcloud() error {
	_, err := exec.Run(exec.Command("which", "gcloud"))
	return err
}
