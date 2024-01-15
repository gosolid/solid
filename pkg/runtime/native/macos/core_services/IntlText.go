//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.IntlText
*/

type IntlText struct {
  TheScriptCode int16 `v8:"theScriptCode"`
  TheLangCode int16 `v8:"theLangCode"`
  TheText [1]byte `v8:"theText"`
}
