package cmd

import (
	"fmt"
	"os"

	"github.com/ros-e/kuro/internal"
	"github.com/ros-e/kuro/internal/checks"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
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
		if !checks.CheckSetup() {
			internal.Error("kuro is not set up, run 'kuro setup' first")
			return fmt.Errorf("")
		}
		name := args[0]
		/*
		* create project folder in /etc/kuro/projects/<name>
		* create config file aswell ^_^ /etc/kuro/projects/<name>/project.yaml
		* n add dat hoe 2 /etc/kuro/config.yaml
		* side note: i could probably remove the {name} param in config.yaml and just make it an array of dirs
		 */
		//ill just move this to a reusable func later
		err := os.MkdirAll(fmt.Sprintf("/etc/kuro/projects/%s", name), 0755)
		if err != nil {
			// this might be fucked because of SilenceErrors ill fix it later :sob:
			return fmt.Errorf("failed to create project folder: %w", err)
		}
		// create config file in project folder
		projectConfig := map[string]interface{}{
			"project_name":        name,
			"project_description": "",
			"services":            map[string]interface{}{},
		}
		projectConfigBytes, err := yaml.Marshal(projectConfig)
		if err != nil {
			return fmt.Errorf("failed to marshal project config: %w", err)
		}
		err = os.WriteFile(fmt.Sprintf("/etc/kuro/projects/%s/project.yaml", name), projectConfigBytes, 0644)
		if err != nil {
			return fmt.Errorf("failed to create config file: %w", err)
		}
		// add project to root/config.yaml
		config, err := os.ReadFile("/etc/kuro/config.yaml")
		if err != nil {
			return fmt.Errorf("failed to read config.yaml: %w", err)
		}
		var rootConfig internal.RootConfig
		err = yaml.Unmarshal(config, &rootConfig)
		if err != nil {
			return fmt.Errorf("failed to unmarshal config.yaml: %w", err)
		}
		if rootConfig.Projects == nil {
			rootConfig.Projects = make(map[string]internal.ProjectEntry)
		}
		rootConfig.Projects[name] = internal.ProjectEntry{
			Name: name,
			Dir:  fmt.Sprintf("/etc/kuro/projects/%s", name),
		}
		configBytes, err := yaml.Marshal(rootConfig)
		if err != nil {
			return fmt.Errorf("failed to marshal config.yaml: %w", err)
		}
		err = os.WriteFile("/etc/kuro/config.yaml", configBytes, 0644)
		if err != nil {
			return fmt.Errorf("failed to write config.yaml: %w", err)
		}
		internal.Success(fmt.Sprintf("Project %q added", name))
		return nil
	},
}

func init() {
	ProjectCmd.AddCommand(projectAddCmd)
}
