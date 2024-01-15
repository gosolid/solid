//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TECEncodingPairs
*/

type TECEncodingPairs struct {
  EncodingPair TECEncodingPairRec `v8:"encodingPair"`
  Flags uint `v8:"flags"`
  Speed uint `v8:"speed"`
}
