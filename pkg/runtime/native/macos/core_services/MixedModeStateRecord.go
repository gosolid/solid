//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.MixedModeStateRecord
*/

type MixedModeStateRecord struct {
  State1 uint `v8:"state1"`
  State2 uint `v8:"state2"`
  State3 uint `v8:"state3"`
  State4 uint `v8:"state4"`
}
