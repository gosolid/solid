//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TECEncodingPairsRec
*/

type TECEncodingPairsRec struct {
  Count uint `v8:"count"`
  EncodingPairs TECEncodingPairs `v8:"encodingPairs"`
}
