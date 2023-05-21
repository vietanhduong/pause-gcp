package version

import (
	"github.com/spf13/cobra"
	"github.com/vietanhduong/pause-gcp/pkg/config/version"
)

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "Print pause-gcp version",
		Run:   func(cmd *cobra.Command, args []string) { version.ShowVersion() },
	}
	return cmd
}
