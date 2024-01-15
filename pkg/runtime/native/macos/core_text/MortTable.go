//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.MortTable
*/

type MortTable struct {
  Version int `v8:"version"`
  NChains uint `v8:"nChains"`
  Chains [1]MortChain `v8:"chains"`
}
