//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KerxKerningPair
*/

type KerxKerningPair struct {
  Left uint16 `v8:"left"`
  Right uint16 `v8:"right"`
}
