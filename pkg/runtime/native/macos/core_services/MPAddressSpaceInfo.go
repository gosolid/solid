//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.MPAddressSpaceInfo
*/

type MPAddressSpaceInfo struct {
  Version uint `v8:"version"`
  ProcessID *OpaqueMPProcessID `v8:"processID"`
  GroupID *OpaqueMPCoherenceID `v8:"groupID"`
  NTasks uint64 `v8:"nTasks"`
  Vsid [16]uint `v8:"vsid"`
}
