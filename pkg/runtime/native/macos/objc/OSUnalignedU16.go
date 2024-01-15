//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc._OSUnalignedU16
*/

type OSUnalignedU16 struct {
  Val uint16 `v8:"val"`
}
