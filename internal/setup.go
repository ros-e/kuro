package internal

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/ros-e/kuro/internal/checks"
)

// might remove verbose ts lk retarded icl :/
func Setup(verbose bool) error {
	var config_location = "/etc/kuro"
	var config_file = "config.yaml"
	// detect if /etc/kuro/ exists and ask for conformation before reinstall
	exists, err := os.Stat(config_location)
	if err == nil && exists.IsDir() {
		if verbose {
			Info("Configuration folder already exists.")
		}
		fmt.Fprint(os.Stdout, "Do you want to reinstall Kuro? (y/n): ")
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))
		if response != "y" {
			return nil
		}
		// delete all configs
		if verbose {
			Info("Deleting existing configuration...")
		}
		if err := os.RemoveAll(config_location); err != nil {
			Error("Failed to delete existing configuration.")
			return fmt.Errorf("failed to delete existing configuration: %w", err)
		}
		if verbose {
			Success("Configuration deleted.")
		}
	}
	// disabled at the moment...
	// no longer disabled(im to lazy to remove that comment ^~^)
	if os.Getuid() != 0 {
		Error("setup must be run as root.")
		return fmt.Errorf("must be run as root")
	}
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " Checking Docker..."
	s.Start()
	dockerOk := checks.CheckDocker()
	s.Stop()
	if !dockerOk {
		Error("Docker not found.")
		if os.Getuid() != 0 {
			Error("Cannot install Docker without root.")
			return errors.New("must be run as root to install Docker")
		}
		Info("Installing Docker...")
		cmd := exec.Command("sh", "-c", "curl -sSL https://get.docker.com | sh")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			Error("Failed to install Docker.")
			return errors.New("docker installation failed")
		}
		if verbose {
			Success("Docker installed.")
		}
	}
	// Create /etc/kuro
	if err := os.MkdirAll(config_location, os.ModePerm); err != nil {
		Error("Failed to create config directory.")
		return fmt.Errorf("failed to create config directory: %w", err)
	}
	if verbose {
		Success("Configuration folder created.")
	}
	// Create config.yaml
	configPath := fmt.Sprintf("%s/%s", config_location, config_file)
	// yes i am aware this is lazy as shit...
	err = os.WriteFile(configPath, []byte("version: \"1.0\"\nprojects: {}\n"), 0644)
	if err != nil {
		Error("Failed to create config.yaml.")
		return fmt.Errorf("failed to create config.yaml: %w", err)
	}
	if verbose {
		Success("config.yaml created.")
	}
	Success("Kuro installed successfully.")
	return nil
}
