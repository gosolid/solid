//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.WritingCode
*/

type WritingCode struct {
  TheScriptCode int16 `v8:"theScriptCode"`
  TheLangCode int16 `v8:"theLangCode"`
}
