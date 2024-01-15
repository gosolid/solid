//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.MPEventInfo
*/

type MPEventInfo struct {
  Version uint `v8:"version"`
  ProcessID *OpaqueMPProcessID `v8:"processID"`
  EventName uint `v8:"eventName"`
  NWaiting uint64 `v8:"nWaiting"`
  WaitingTaskID *OpaqueMPTaskID `v8:"waitingTaskID"`
  Events uint `v8:"events"`
}
