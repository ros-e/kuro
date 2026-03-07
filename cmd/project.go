package cmd

import (
	"fmt"

	"github.com/ros-e/kuro/internal"
	"github.com/spf13/cobra"
)

var ProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
}

var projectAddCmd = &cobra.Command{
	Use:   "add <name>",
	Short: "Add a project to Kuro",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		internal.Success(fmt.Sprintf("Project %q added", name))
		return nil
	},
}

func init() {
	ProjectCmd.AddCommand(projectAddCmd)
}
