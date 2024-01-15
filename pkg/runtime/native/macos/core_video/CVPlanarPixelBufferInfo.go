//js:package native/macos/core-video
package core_video

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreVideo.CVPlanarPixelBufferInfo
*/

type CVPlanarPixelBufferInfo struct {
  ComponentInfo [1]CVPlanarComponentInfo `v8:"componentInfo"`
}
