//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.FSCatalogInfo
*/

type FSCatalogInfo struct {
  NodeFlags uint16 `v8:"nodeFlags"`
  Volume int16 `v8:"volume"`
  ParentDirID uint `v8:"parentDirID"`
  NodeID uint `v8:"nodeID"`
  SharingFlags byte `v8:"sharingFlags"`
  UserPrivileges byte `v8:"userPrivileges"`
  Reserved1 byte `v8:"reserved1"`
  Reserved2 byte `v8:"reserved2"`
  CreateDate UTCDateTime `v8:"createDate"`
  ContentModDate UTCDateTime `v8:"contentModDate"`
  AttributeModDate UTCDateTime `v8:"attributeModDate"`
  AccessDate UTCDateTime `v8:"accessDate"`
  BackupDate UTCDateTime `v8:"backupDate"`
  Permissions FSPermissionInfo `v8:"permissions"`
  FinderInfo [16]byte `v8:"finderInfo"`
  ExtFinderInfo [16]byte `v8:"extFinderInfo"`
  DataLogicalSize uint64 `v8:"dataLogicalSize"`
  DataPhysicalSize uint64 `v8:"dataPhysicalSize"`
  RsrcLogicalSize uint64 `v8:"rsrcLogicalSize"`
  RsrcPhysicalSize uint64 `v8:"rsrcPhysicalSize"`
  Valence uint `v8:"valence"`
  TextEncodingHint uint `v8:"textEncodingHint"`
}
