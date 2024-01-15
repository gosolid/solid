//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.SCSITaskDeviceInterface
*/

type SCSITaskDeviceInterface struct {
  Reserved void `v8:"reserved"`
  QueryInterface func(...any) (any, error) `v8:"queryInterface"`
  AddRef func(...any) (any, error) `v8:"addRef"`
  Release func(...any) (any, error) `v8:"release"`
  Version uint16 `v8:"version"`
  Revision uint16 `v8:"revision"`
  IsExclusiveAccessAvailable func(...any) (any, error) `v8:"isExclusiveAccessAvailable"`
  AddCallbackDispatcherToRunLoop func(...any) (any, error) `v8:"addCallbackDispatcherToRunLoop"`
  RemoveCallbackDispatcherFromRunLoop func(...any) (any, error) `v8:"removeCallbackDispatcherFromRunLoop"`
  ObtainExclusiveAccess func(...any) (any, error) `v8:"obtainExclusiveAccess"`
  ReleaseExclusiveAccess func(...any) (any, error) `v8:"releaseExclusiveAccess"`
  CreateSCSITask func(...any) (any, error) `v8:"createSCSITask"`
}
