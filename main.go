package main

import (
	"os"

	"github.com/ros-e/kuro/cmd"
	"github.com/ros-e/kuro/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "kuro",
	Short:   "Self-Hosted PaaS alternative to Vercel and Heroku",
	Version: internal.Version,
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
	rootCmd.AddCommand(cmd.SetupCmd)
	rootCmd.AddCommand(cmd.ServiceCmd)
	rootCmd.AddCommand(cmd.ProjectCmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
