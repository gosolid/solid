//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.MPSemaphoreInfo
*/

type MPSemaphoreInfo struct {
  Version uint `v8:"version"`
  ProcessID *OpaqueMPProcessID `v8:"processID"`
  SemaphoreName uint `v8:"semaphoreName"`
  NWaiting uint64 `v8:"nWaiting"`
  WaitingTaskID *OpaqueMPTaskID `v8:"waitingTaskID"`
  Maximum uint64 `v8:"maximum"`
  Count uint64 `v8:"count"`
}
