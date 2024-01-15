//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.BigEndianUnsignedShort
*/

type BigEndianUnsignedShort struct {
  BigEndianValue uint16 `v8:"bigEndianValue"`
}
