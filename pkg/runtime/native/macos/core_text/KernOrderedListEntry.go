//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KernOrderedListEntry
*/

type KernOrderedListEntry struct {
  Pair KernKerningPair `v8:"pair"`
  Value int16 `v8:"value"`
}
