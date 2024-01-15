//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TECConversionInfo
*/

type TECConversionInfo struct {
  SourceEncoding uint `v8:"sourceEncoding"`
  DestinationEncoding uint `v8:"destinationEncoding"`
  Reserved1 uint16 `v8:"reserved1"`
  Reserved2 uint16 `v8:"reserved2"`
}
