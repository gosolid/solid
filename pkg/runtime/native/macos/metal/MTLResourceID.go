//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLResourceID
*/

type MTLResourceID struct {
  Impl uint64 `v8:"impl"`
}
