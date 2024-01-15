//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.BslnFormat1Part
*/

type BslnFormat1Part struct {
  Deltas [32]int16 `v8:"deltas"`
  MappingData SFNTLookupTable `v8:"mappingData"`
}
