package version

import "runtime"

func GetVersion() string {
	return runtime.Version()
}
