//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.PEFLoaderRelocationHeader
*/

type PEFLoaderRelocationHeader struct {
  SectionIndex uint16 `v8:"sectionIndex"`
  ReservedA uint16 `v8:"reservedA"`
  RelocCount uint `v8:"relocCount"`
  FirstRelocOffset uint `v8:"firstRelocOffset"`
}
