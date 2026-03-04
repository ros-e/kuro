package cmd

import (
	"fmt"

	"github.com/ros-e/kuro/internal/checks"
	"github.com/spf13/cobra"
)

var SetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Configure and initialize Kuro",
	RunE: func(cmd *cobra.Command, args []string) error {
		if checks.CheckDocker() {
			fmt.Println("Docker executable found in PATH.")
		} else {
			fmt.Println("Docker executable not found in PATH. Docker may not be installed.")
		}
		return nil
	},
}
