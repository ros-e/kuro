package cmd

import (
	"github.com/ros-e/kuro/internal"
	"github.com/spf13/cobra"
)

var SetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Configure and initialize Kuro",
	RunE: func(cmd *cobra.Command, args []string) error {
		verbose, _ := cmd.Flags().GetBool("verbose")
		internal.Setup(verbose)
		return nil
	},
}
