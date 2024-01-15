//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc._OSUnalignedU64
*/

type OSUnalignedU64 struct {
  Val uint64 `v8:"val"`
}
