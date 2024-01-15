//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.NumberParts
*/

type NumberParts struct {
  Version int16 `v8:"version"`
  Data void `v8:"data"`
  PePlus WideCharArr `v8:"pePlus"`
  PeMinus WideCharArr `v8:"peMinus"`
  PeMinusPlus WideCharArr `v8:"peMinusPlus"`
  AltNumTable WideCharArr `v8:"altNumTable"`
  Reserved [20]byte `v8:"reserved"`
}
