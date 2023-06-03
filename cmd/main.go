package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/vietanhduong/pause-gcp/cmd/gke"
	"github.com/vietanhduong/pause-gcp/cmd/sql"
	"github.com/vietanhduong/pause-gcp/cmd/version"
	"github.com/vietanhduong/pause-gcp/cmd/vm"
	"github.com/vietanhduong/pause-gcp/pkg/utils/exec"
)

func newRootCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "pause-gcp",
	}
	cmd.AddCommand(version.NewCommand())
	cmd.AddCommand(gke.NewCommand())
	cmd.AddCommand(vm.NewCommand())
	cmd.AddCommand(sql.NewCommand())
	return cmd
}

func main() {
	if err := testGcloud(); err != nil {
		log.Fatalln(err)
	}
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}

func testGcloud() error {
	_, err := exec.Run(exec.Command("which", "gcloud"))
	return err
}
