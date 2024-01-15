//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.FSVolumeInfo
*/

type FSVolumeInfo struct {
  CreateDate UTCDateTime `v8:"createDate"`
  ModifyDate UTCDateTime `v8:"modifyDate"`
  BackupDate UTCDateTime `v8:"backupDate"`
  CheckedDate UTCDateTime `v8:"checkedDate"`
  FileCount uint `v8:"fileCount"`
  FolderCount uint `v8:"folderCount"`
  TotalBytes uint64 `v8:"totalBytes"`
  FreeBytes uint64 `v8:"freeBytes"`
  BlockSize uint `v8:"blockSize"`
  TotalBlocks uint `v8:"totalBlocks"`
  FreeBlocks uint `v8:"freeBlocks"`
  NextAllocation uint `v8:"nextAllocation"`
  RsrcClumpSize uint `v8:"rsrcClumpSize"`
  DataClumpSize uint `v8:"dataClumpSize"`
  NextCatalogID uint `v8:"nextCatalogID"`
  FinderInfo [32]byte `v8:"finderInfo"`
  Flags uint16 `v8:"flags"`
  FilesystemID uint16 `v8:"filesystemID"`
  Signature uint16 `v8:"signature"`
  DriveNumber uint16 `v8:"driveNumber"`
  DriverRefNum int `v8:"driverRefNum"`
}
