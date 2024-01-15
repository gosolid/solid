//js:package native/macos/core-video
package core_video

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreVideo.CVTimeStamp
*/

type CVTimeStamp struct {
  Version uint `v8:"version"`
  VideoTimeScale int `v8:"videoTimeScale"`
  VideoTime int64 `v8:"videoTime"`
  HostTime uint64 `v8:"hostTime"`
  RateScalar float64 `v8:"rateScalar"`
  VideoRefreshPeriod int64 `v8:"videoRefreshPeriod"`
  SmpteTime CVSMPTETime `v8:"smpteTime"`
  Flags uint64 `v8:"flags"`
  Reserved uint64 `v8:"reserved"`
}
