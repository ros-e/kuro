package checks

import "os/exec"

func CheckDocker() bool {
	_, err := exec.LookPath("docker")
	return err == nil
}
