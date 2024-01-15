//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc._OSUnalignedU32
*/

type OSUnalignedU32 struct {
  Val uint `v8:"val"`
}
