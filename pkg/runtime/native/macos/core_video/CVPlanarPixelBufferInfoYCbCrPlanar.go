//js:package native/macos/core-video
package core_video

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreVideo.CVPlanarPixelBufferInfo_YCbCrPlanar
*/

type CVPlanarPixelBufferInfoYCbCrPlanar struct {
  ComponentInfoY CVPlanarComponentInfo `v8:"componentInfoY"`
  ComponentInfoCb CVPlanarComponentInfo `v8:"componentInfoCb"`
  ComponentInfoCr CVPlanarComponentInfo `v8:"componentInfoCr"`
}
