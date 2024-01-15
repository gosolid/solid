//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.PEFSectionHeader
*/

type PEFSectionHeader struct {
  NameOffset int `v8:"nameOffset"`
  DefaultAddress uint `v8:"defaultAddress"`
  TotalLength uint `v8:"totalLength"`
  UnpackedLength uint `v8:"unpackedLength"`
  ContainerLength uint `v8:"containerLength"`
  ContainerOffset uint `v8:"containerOffset"`
  SectionKind byte `v8:"sectionKind"`
  ShareKind byte `v8:"shareKind"`
  Alignment byte `v8:"alignment"`
  ReservedA byte `v8:"reservedA"`
}
