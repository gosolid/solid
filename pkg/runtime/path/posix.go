//js:package path

package path

//go:generate go run github.com/grexie/isolates/codegen

//go:generate go run github.com/grexie/isolates/codegen

type Import struct{}

//js:method
//js:export toNamespacedPath
func ToNamespacedPath(path string) string {
	return path
}
