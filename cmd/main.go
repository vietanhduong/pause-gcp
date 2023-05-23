package main

import (
	"github.com/spf13/cobra"
	"github.com/vietanhduong/pause-gcp/cmd/pause"
	"github.com/vietanhduong/pause-gcp/cmd/unpause"
	"github.com/vietanhduong/pause-gcp/cmd/version"
	"github.com/vietanhduong/pause-gcp/pkg/utils/exec"
	"log"
	"os"
)

func newRootCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "pause-gcp",
	}
	cmd.AddCommand(version.NewCommand())
	cmd.AddCommand(pause.NewCommand())
	cmd.AddCommand(unpause.NewCommand())
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
