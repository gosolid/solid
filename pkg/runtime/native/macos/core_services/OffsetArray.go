//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.OffsetArray
*/

type OffsetArray struct {
  FNumOfOffsets int16 `v8:"fNumOfOffsets"`
  FOffset [1]int `v8:"fOffset"`
}
