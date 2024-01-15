//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CFAllocatorContext
*/

type CFAllocatorContext struct {
  Version int64 `v8:"version"`
  Info void `v8:"info"`
  Retain *func(...any) (any, error) `v8:"retain"`
  Release *func(...any) (any, error) `v8:"release"`
  CopyDescription *func(...any) (any, error) `v8:"copyDescription"`
  Allocate *func(...any) (any, error) `v8:"allocate"`
  Reallocate *func(...any) (any, error) `v8:"reallocate"`
  Deallocate *func(...any) (any, error) `v8:"deallocate"`
  PreferredSize *func(...any) (any, error) `v8:"preferredSize"`
}
