//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ComponentParameters
*/

type ComponentParameters struct {
  Flags byte `v8:"flags"`
  ParamSize byte `v8:"paramSize"`
  What int16 `v8:"what"`
  Padding uint `v8:"padding"`
  Params [1]int64 `v8:"params"`
}
