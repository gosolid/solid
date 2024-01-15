//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TECEncodingPairRec
*/

type TECEncodingPairRec struct {
  Source TextEncodingRec `v8:"source"`
  Dest TextEncodingRec `v8:"dest"`
}
