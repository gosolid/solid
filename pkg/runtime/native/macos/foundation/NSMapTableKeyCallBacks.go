//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Foundation.NSMapTableKeyCallBacks
*/

type NSMapTableKeyCallBacks struct {
  Hash func(...any) (any, error) `v8:"hash"`
  IsEqual func(...any) (any, error) `v8:"isEqual"`
  Retain func(...any) (any, error) `v8:"retain"`
  Release func(...any) (any, error) `v8:"release"`
  Describe func(...any) (any, error) `v8:"describe"`
  NotAKeyMarker void `v8:"notAKeyMarker"`
}
