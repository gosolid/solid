//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TextRange
*/

type TextRange struct {
  FStart int `v8:"fStart"`
  FEnd int `v8:"fEnd"`
  FHiliteStyle int16 `v8:"fHiliteStyle"`
}
