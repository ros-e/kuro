package cmd

import (
	"fmt"

	"github.com/ros-e/kuro/internal"
	"github.com/spf13/cobra"
)

var ServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Manage services inside a Kuro project",
}

var serviceAddCmd = &cobra.Command{
	Use:   "add <name>",
	Short: "Add a service to a project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		project, _ := cmd.Flags().GetString("project")

		if project == "" {
			return fmt.Errorf("--project is required")
		}

		internal.Success(fmt.Sprintf("Service %q added to project %q", name, project))
		return nil
	},
}

func init() {
	serviceAddCmd.Flags().StringP("project", "p", "", "Project to add the service to (required)")
	ServiceCmd.AddCommand(serviceAddCmd)
}
