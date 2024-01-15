//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TextEncodingRun
*/

type TextEncodingRun struct {
  Offset uint64 `v8:"offset"`
  TextEncoding uint `v8:"textEncoding"`
}
