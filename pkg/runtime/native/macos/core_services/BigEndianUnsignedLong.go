//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.BigEndianUnsignedLong
*/

type BigEndianUnsignedLong struct {
  BigEndianValue uint64 `v8:"bigEndianValue"`
}
