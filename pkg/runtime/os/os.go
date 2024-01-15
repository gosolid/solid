//js:package os

package os

//go:generate go run github.com/grexie/isolates/codegen

import (
	"runtime"
)

//go:generate go run github.com/grexie/isolates/codegen

type Import struct{}

//js:method
//js:export platform
func Platform() string {
	switch runtime.GOOS {
	case "windows":
		return "win32"
	default:
		return runtime.GOOS
	}
}

//js:method
//js:export type
func Type() (string, error) {
	return os_type()
}
