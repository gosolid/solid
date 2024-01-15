//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.FSForkInfo
*/

type FSForkInfo struct {
  Flags byte `v8:"flags"`
  Permissions byte `v8:"permissions"`
  Volume int16 `v8:"volume"`
  Reserved2 uint `v8:"reserved2"`
  NodeID uint `v8:"nodeID"`
  ForkID uint `v8:"forkID"`
  CurrentPosition uint64 `v8:"currentPosition"`
  LogicalEOF uint64 `v8:"logicalEOF"`
  PhysicalEOF uint64 `v8:"physicalEOF"`
  Process uint64 `v8:"process"`
}
