//js:package tty

package tty

//go:generate go run github.com/grexie/isolates/codegen

//go:generate go run github.com/grexie/isolates/codegen

//js:method isatty
//js:export isatty
func IsTTY(fd int) (bool, error) {
	return true, nil
}
