//js:package native/macos/core-video
package core_video

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreVideo.CVFillExtendedPixelsCallBackData
*/

type CVFillExtendedPixelsCallBackData struct {
  Version int64 `v8:"version"`
  FillCallBack *func(...any) (any, error) `v8:"fillCallBack"`
  RefCon void `v8:"refCon"`
}
