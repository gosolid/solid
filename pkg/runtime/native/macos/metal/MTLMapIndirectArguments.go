//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLMapIndirectArguments
*/

type MTLMapIndirectArguments struct {
  RegionOriginX uint `v8:"regionOriginX"`
  RegionOriginY uint `v8:"regionOriginY"`
  RegionOriginZ uint `v8:"regionOriginZ"`
  RegionSizeWidth uint `v8:"regionSizeWidth"`
  RegionSizeHeight uint `v8:"regionSizeHeight"`
  RegionSizeDepth uint `v8:"regionSizeDepth"`
  MipMapLevel uint `v8:"mipMapLevel"`
  SliceId uint `v8:"sliceId"`
}
