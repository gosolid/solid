//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TextEncodingRec
*/

type TextEncodingRec struct {
  Base uint `v8:"base"`
  Variant uint `v8:"variant"`
  Format uint `v8:"format"`
}
