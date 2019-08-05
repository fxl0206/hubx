package root

import (
	"github.com/spf13/cobra"
)

var(
	RootCmd = &cobra.Command{
		Use:          "hubxs",
		Short:        "hubxs",
		Long:         "hubx-server",
		SilenceUsage: true,
	}
)