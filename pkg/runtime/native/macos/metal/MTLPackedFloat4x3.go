//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLPackedFloat4x3
*/

type MTLPackedFloat4x3 struct {
  Columns [4]MTLPackedFloat3 `v8:"columns"`
}
