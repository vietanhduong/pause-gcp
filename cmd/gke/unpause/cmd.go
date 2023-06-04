package unpause

import (
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	apis "github.com/vietanhduong/pause-gcp/apis/v1"
	"github.com/vietanhduong/pause-gcp/pkg/utils/exec"
	"github.com/vietanhduong/pause-gcp/pkg/utils/protoutil"
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
		Example: `
# STATE_FILE from local
$ pause-gcp gke unpause ./gke-states/gke_develop_asia-southeast1_dev-cluster.state.json

# STATE_FILE from a gcs bucket
$ pause-gcp gke unpause gs://bucket/path/json_file.json --rm
`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.Errorf("STATE_FILE is missing")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var b []byte
			var err error
			// if the input path starts with 'gs://', that means the state is stored at GCS
			// and we can use gsutil cat to retrieve the file content
			if strings.HasPrefix(strings.ToLower(args[0]), "gs://") {
				var raw string
				if raw, err = exec.Run(exec.Command("gsutil", "cat", args[0])); err != nil {
					return err
				}
				b = []byte(raw)
			} else {
				if b, err = os.ReadFile(args[0]); err != nil {
					return err
				}
			}
			log.Printf("INFO: retrieve cluster state completed!\n")
			var cluster apis.Cluster
			if err = protoutil.Unmarshal(b, &cluster); err != nil {
				return err
			}
			runCfg.cluster = &cluster
			if err = run(runCfg); err != nil {
				return err
			}
			// remote the input state file if the --rm flag exists
			if rm {
				// remove the state file by gsutil if the input path start with gs://
				if strings.HasPrefix(strings.ToLower(args[0]), "gs://") {
					_, err = exec.Run(exec.Command("gsutil", "rm", args[0]))
					return err
				} else {
					return os.Remove(args[0])
				}
			}
			return nil
		},
	}

	cmd.Flags().BoolVar(&rm, "rm", false, "Remove the cluster state after complete")

	return cmd
}
