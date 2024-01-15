//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.decform
*/

type Decform struct {
  Style byte `v8:"style"`
  Unused byte `v8:"unused"`
  Digits int16 `v8:"digits"`
}
