//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.WideCharArr
*/

type WideCharArr struct {
  Size int16 `v8:"size"`
  Data void `v8:"data"`
}
