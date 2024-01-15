//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.CTRunDelegateCallbacks
*/

type CTRunDelegateCallbacks struct {
  Version int64 `v8:"version"`
  Dealloc *func(...any) (any, error) `v8:"dealloc"`
  GetAscent *func(...any) (any, error) `v8:"getAscent"`
  GetDescent *func(...any) (any, error) `v8:"getDescent"`
  GetWidth *func(...any) (any, error) `v8:"getWidth"`
}
