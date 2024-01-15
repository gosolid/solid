//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.siginfo_t
*/

type SiginfoT struct {
  SiSigno int `v8:"siSigno"`
  SiErrno int `v8:"siErrno"`
  SiCode int `v8:"siCode"`
  SiPid int `v8:"siPid"`
  SiUid uint `v8:"siUid"`
  SiStatus int `v8:"siStatus"`
  SiAddr void `v8:"siAddr"`
  SiValue void `v8:"siValue"`
  SiBand int64 `v8:"siBand"`
  Pad [7]uint64 `v8:"pad"`
}
