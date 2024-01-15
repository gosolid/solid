//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.UCKeyStateEntryRange
*/

type UCKeyStateEntryRange struct {
  CurStateStart uint16 `v8:"curStateStart"`
  CurStateRange byte `v8:"curStateRange"`
  DeltaMultiplier byte `v8:"deltaMultiplier"`
  CharData uint16 `v8:"charData"`
  NextState uint16 `v8:"nextState"`
}
