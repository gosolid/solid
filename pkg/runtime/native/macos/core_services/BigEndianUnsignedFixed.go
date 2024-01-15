//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.BigEndianUnsignedFixed
*/

type BigEndianUnsignedFixed struct {
  BigEndianValue uint `v8:"bigEndianValue"`
}
