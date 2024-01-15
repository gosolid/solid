//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.proc_rlimit_control_wakeupmon
*/

type ProcRlimitControlWakeupmon struct {
  WmFlags uint `v8:"wmFlags"`
  WmRate int `v8:"wmRate"`
}
