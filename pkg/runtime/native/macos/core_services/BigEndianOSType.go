//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.BigEndianOSType
*/

type BigEndianOSType struct {
  BigEndianValue uint `v8:"bigEndianValue"`
}
