//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.objc_object
*/

type ObjcObject struct {
  Isa *any `v8:"isa"`
}
