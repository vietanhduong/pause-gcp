package utils

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"strings"
)

func RequiredFlags(cmd *cobra.Command, flag ...string) error {
	fs := make(map[string]string)
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		fs[flag.Name] = flag.Value.String()
	})

	var missing []string

	for _, f := range flag {
		v, found := fs[f]
		if !found || v == "" {
			missing = append(missing, fmt.Sprintf("'%s'", f))
		}
	}
	if len(missing) > 0 {
		miss := strings.Join(missing, ", ")
		return fmt.Errorf("required flag(s) %s not set", miss)
	}
	return nil
}
