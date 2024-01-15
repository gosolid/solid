//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.UCKeyToCharTableIndex
*/

type UCKeyToCharTableIndex struct {
  KeyToCharTableIndexFormat uint16 `v8:"keyToCharTableIndexFormat"`
  KeyToCharTableSize uint16 `v8:"keyToCharTableSize"`
  KeyToCharTableCount uint `v8:"keyToCharTableCount"`
  KeyToCharTableOffsets [1]uint `v8:"keyToCharTableOffsets"`
}
