package cmd

import (
	"github.com/ros-e/kuro/internal"
	"github.com/ros-e/kuro/internal/checks"
	"github.com/spf13/cobra"
)

var SetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Configure and initialize Kuro",
	RunE: func(cmd *cobra.Command, args []string) error {
		if checks.CheckDocker() {
			internal.Success("Docker executable found in PATH.")
		} else {
			internal.Error("Docker executable not found in PATH. Docker may not be installed.")
		}
		return nil
	},
}
