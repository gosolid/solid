//js:package native/macos/core-video
package core_video

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreVideo.CVPlanarPixelBufferInfo_YCbCrBiPlanar
*/

type CVPlanarPixelBufferInfoYCbCrBiPlanar struct {
  ComponentInfoY CVPlanarComponentInfo `v8:"componentInfoY"`
  ComponentInfoCbCr CVPlanarComponentInfo `v8:"componentInfoCbCr"`
}
