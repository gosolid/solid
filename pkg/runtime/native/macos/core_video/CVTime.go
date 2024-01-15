//js:package native/macos/core-video
package core_video

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreVideo.CVTime
*/

type CVTime struct {
  TimeValue int64 `v8:"timeValue"`
  TimeScale int `v8:"timeScale"`
  Flags int `v8:"flags"`
}
