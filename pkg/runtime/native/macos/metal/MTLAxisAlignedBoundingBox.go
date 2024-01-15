//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLAxisAlignedBoundingBox
*/

type MTLAxisAlignedBoundingBox struct {
  Min MTLPackedFloat3 `v8:"min"`
  Max MTLPackedFloat3 `v8:"max"`
}
