//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.BigEndianShort
*/

type BigEndianShort struct {
  BigEndianValue int16 `v8:"bigEndianValue"`
}
