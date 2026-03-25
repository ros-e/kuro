package checks

import "os"

/*  js check if kuro has been setup */
func CheckSetup() bool {
	_, err := os.Stat("/etc/kuro/config.yaml")
	return err == nil
}
