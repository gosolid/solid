//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TECEncodingsListRec
*/

type TECEncodingsListRec struct {
  Count uint `v8:"count"`
  Encodings TextEncodingRec `v8:"encodings"`
}
