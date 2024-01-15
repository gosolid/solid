//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TECInternetNameRec
*/

type TECInternetNameRec struct {
  Offset uint `v8:"offset"`
  SearchEncoding TextEncodingRec `v8:"searchEncoding"`
  EncodingNameLength byte `v8:"encodingNameLength"`
  EncodingName [1]byte `v8:"encodingName"`
}
