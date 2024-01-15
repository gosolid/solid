//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLViewport
*/

type MTLViewport struct {
  OriginX float64 `v8:"originX"`
  OriginY float64 `v8:"originY"`
  Width float64 `v8:"width"`
  Height float64 `v8:"height"`
  Znear float64 `v8:"znear"`
  Zfar float64 `v8:"zfar"`
}
