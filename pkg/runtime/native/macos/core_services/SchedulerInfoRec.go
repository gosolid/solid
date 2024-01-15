//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.SchedulerInfoRec
*/

type SchedulerInfoRec struct {
  InfoRecSize uint `v8:"infoRecSize"`
  CurrentThreadID uint64 `v8:"currentThreadID"`
  SuggestedThreadID uint64 `v8:"suggestedThreadID"`
  InterruptedCoopThreadID uint64 `v8:"interruptedCoopThreadID"`
}
