//js:package native/macos/core-video
package core_video

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreVideo.CVSMPTETime
*/

type CVSMPTETime struct {
  Subframes int16 `v8:"subframes"`
  SubframeDivisor int16 `v8:"subframeDivisor"`
  Counter uint `v8:"counter"`
  Type uint `v8:"type"`
  Flags uint `v8:"flags"`
  Hours int16 `v8:"hours"`
  Minutes int16 `v8:"minutes"`
  Seconds int16 `v8:"seconds"`
  Frames int16 `v8:"frames"`
}
