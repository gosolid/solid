//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.RoutingResourceEntry
*/

type RoutingResourceEntry struct {
  Creator uint `v8:"creator"`
  FileType uint `v8:"fileType"`
  TargetFolder uint `v8:"targetFolder"`
  DestinationFolder uint `v8:"destinationFolder"`
  ReservedField uint `v8:"reservedField"`
}
