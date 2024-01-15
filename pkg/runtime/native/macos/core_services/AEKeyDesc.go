//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.AEKeyDesc
*/

type AEKeyDesc struct {
  DescKey uint `v8:"descKey"`
  DescContent AEDesc `v8:"descContent"`
}
