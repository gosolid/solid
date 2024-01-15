//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Foundation.NSMapTableValueCallBacks
*/

type NSMapTableValueCallBacks struct {
  Retain func(...any) (any, error) `v8:"retain"`
  Release func(...any) (any, error) `v8:"release"`
  Describe func(...any) (any, error) `v8:"describe"`
}
