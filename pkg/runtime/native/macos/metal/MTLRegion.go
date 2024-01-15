//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLRegion
*/

type MTLRegion struct {
  Origin MTLOrigin `v8:"origin"`
  Size MTLSize `v8:"size"`
}
