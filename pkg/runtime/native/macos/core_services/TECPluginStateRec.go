//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TECPluginStateRec
*/

type TECPluginStateRec struct {
  State1 byte `v8:"state1"`
  State2 byte `v8:"state2"`
  State3 byte `v8:"state3"`
  State4 byte `v8:"state4"`
  LongState1 uint `v8:"longState1"`
  LongState2 uint `v8:"longState2"`
  LongState3 uint `v8:"longState3"`
  LongState4 uint `v8:"longState4"`
}
