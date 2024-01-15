//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.MorxTable
*/

type MorxTable struct {
  Version int `v8:"version"`
  NChains uint `v8:"nChains"`
  Chains [1]MorxChain `v8:"chains"`
}
