//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLSize
*/

type MTLSize struct {
  Width uint64 `v8:"width"`
  Height uint64 `v8:"height"`
  Depth uint64 `v8:"depth"`
}
