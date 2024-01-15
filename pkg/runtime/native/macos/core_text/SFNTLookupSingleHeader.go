//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.SFNTLookupSingleHeader
*/

type SFNTLookupSingleHeader struct {
  BinSearch SFNTLookupBinarySearchHeader `v8:"binSearch"`
  Entries [1]SFNTLookupSingle `v8:"entries"`
}
