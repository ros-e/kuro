package internal

import (
	"fmt"
	"time"
)

/*
 * Usage
 *
 * 	import (
 * 		"github.com/ros-e/kuro/internal"
 * 	)
 *
 * 	func main() {
 *		internal.Info("Hello World!")
 * 	}
 *
 */

const (
	reset  = "\033[0m"
	green  = "\033[32m"
	red    = "\033[31m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
)

func timestamp() string {
	return time.Now().Format("15:04:05.000")
}

func Success(msg string) {
	fmt.Printf("[%s] %sSUCCESS:%s %s\n", timestamp(), green, reset, msg)
}

func Info(msg string) {
	fmt.Printf("[%s] %sINFO:%s %s\n", timestamp(), cyan, reset, msg)
}

func Error(msg string) {
	fmt.Printf("[%s] %sERROR:%s %s\n", timestamp(), red, reset, msg)
}

func Warn(msg string) {
	fmt.Printf("[%s] %sWARN:%s %s\n", timestamp(), yellow, reset, msg)
}

func Debug(msg string) {
	fmt.Printf("[%s] %sDEBUG:%s %s\n", timestamp(), green, reset, msg)
}
