//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.BigEndianFixed
*/

type BigEndianFixed struct {
  BigEndianValue int `v8:"bigEndianValue"`
}
