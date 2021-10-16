package osutils

import "runtime"

// IdentifyOS returns the OS version (windows, linux, darwin)
func IdentifyOS() string {
	return runtime.GOOS
}
