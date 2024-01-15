//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.BigEndianUInt32
*/

type BigEndianUInt32 struct {
  BigEndianValue uint `v8:"bigEndianValue"`
}
