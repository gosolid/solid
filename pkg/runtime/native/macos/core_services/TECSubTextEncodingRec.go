//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TECSubTextEncodingRec
*/

type TECSubTextEncodingRec struct {
  Offset uint `v8:"offset"`
  SearchEncoding TextEncodingRec `v8:"searchEncoding"`
  Count uint `v8:"count"`
  SubEncodings TextEncodingRec `v8:"subEncodings"`
}
