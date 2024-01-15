//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.MPCriticalRegionInfo
*/

type MPCriticalRegionInfo struct {
  Version uint `v8:"version"`
  ProcessID *OpaqueMPProcessID `v8:"processID"`
  RegionName uint `v8:"regionName"`
  NWaiting uint64 `v8:"nWaiting"`
  WaitingTaskID *OpaqueMPTaskID `v8:"waitingTaskID"`
  OwningTask *OpaqueMPTaskID `v8:"owningTask"`
  Count uint64 `v8:"count"`
}
