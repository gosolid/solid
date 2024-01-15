//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.XLibContainerHeader
*/

type XLibContainerHeader struct {
  Tag1 uint `v8:"tag1"`
  Tag2 uint `v8:"tag2"`
  CurrentFormat uint `v8:"currentFormat"`
  ContainerStringsOffset uint `v8:"containerStringsOffset"`
  ExportHashOffset uint `v8:"exportHashOffset"`
  ExportKeyOffset uint `v8:"exportKeyOffset"`
  ExportSymbolOffset uint `v8:"exportSymbolOffset"`
  ExportNamesOffset uint `v8:"exportNamesOffset"`
  ExportHashTablePower uint `v8:"exportHashTablePower"`
  ExportedSymbolCount uint `v8:"exportedSymbolCount"`
  FragNameOffset uint `v8:"fragNameOffset"`
  FragNameLength uint `v8:"fragNameLength"`
  DylibPathOffset uint `v8:"dylibPathOffset"`
  DylibPathLength uint `v8:"dylibPathLength"`
  CpuFamily uint `v8:"cpuFamily"`
  CpuModel uint `v8:"cpuModel"`
  DateTimeStamp uint `v8:"dateTimeStamp"`
  CurrentVersion uint `v8:"currentVersion"`
  OldDefVersion uint `v8:"oldDefVersion"`
  OldImpVersion uint `v8:"oldImpVersion"`
}
