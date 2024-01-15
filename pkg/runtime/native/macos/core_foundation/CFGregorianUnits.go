//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CFGregorianUnits
*/

type CFGregorianUnits struct {
  Years int `v8:"years"`
  Months int `v8:"months"`
  Days int `v8:"days"`
  Hours int `v8:"hours"`
  Minutes int `v8:"minutes"`
  Seconds float64 `v8:"seconds"`
}
