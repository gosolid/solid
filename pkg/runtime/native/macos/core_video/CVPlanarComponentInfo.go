//js:package native/macos/core-video
package core_video

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreVideo.CVPlanarComponentInfo
*/

type CVPlanarComponentInfo struct {
  Offset int `v8:"offset"`
  RowBytes uint `v8:"rowBytes"`
}
