package sql

import (
	"github.com/spf13/cobra"
	"github.com/vietanhduong/pause-gcp/cmd/sql/pause"
	"github.com/vietanhduong/pause-gcp/cmd/sql/unpause"
)

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "sql",
		Short: "Manage Cloud SQL resource",
		Long:  `The sql command lets you can mange your Cloud SQL instances like pause/unpause`,
	}

	cmd.AddCommand(pause.NewCommand())
	cmd.AddCommand(unpause.NewCommand())
	return cmd
}
