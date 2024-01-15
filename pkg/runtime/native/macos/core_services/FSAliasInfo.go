//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.FSAliasInfo
*/

type FSAliasInfo struct {
  VolumeCreateDate UTCDateTime `v8:"volumeCreateDate"`
  TargetCreateDate UTCDateTime `v8:"targetCreateDate"`
  FileType uint `v8:"fileType"`
  FileCreator uint `v8:"fileCreator"`
  ParentDirID uint `v8:"parentDirID"`
  NodeID uint `v8:"nodeID"`
  FilesystemID uint16 `v8:"filesystemID"`
  Signature uint16 `v8:"signature"`
  VolumeIsBootVolume byte `v8:"volumeIsBootVolume"`
  VolumeIsAutomounted byte `v8:"volumeIsAutomounted"`
  VolumeIsEjectable byte `v8:"volumeIsEjectable"`
  VolumeHasPersistentFileIDs byte `v8:"volumeHasPersistentFileIDs"`
  IsDirectory byte `v8:"isDirectory"`
}
