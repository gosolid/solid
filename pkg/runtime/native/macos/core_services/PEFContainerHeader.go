//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.PEFContainerHeader
*/

type PEFContainerHeader struct {
  Tag1 uint `v8:"tag1"`
  Tag2 uint `v8:"tag2"`
  Architecture uint `v8:"architecture"`
  FormatVersion uint `v8:"formatVersion"`
  DateTimeStamp uint `v8:"dateTimeStamp"`
  OldDefVersion uint `v8:"oldDefVersion"`
  OldImpVersion uint `v8:"oldImpVersion"`
  CurrentVersion uint `v8:"currentVersion"`
  SectionCount uint16 `v8:"sectionCount"`
  InstSectionCount uint16 `v8:"instSectionCount"`
  ReservedA uint `v8:"reservedA"`
}
