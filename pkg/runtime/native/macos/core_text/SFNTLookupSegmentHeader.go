//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.SFNTLookupSegmentHeader
*/

type SFNTLookupSegmentHeader struct {
  BinSearch SFNTLookupBinarySearchHeader `v8:"binSearch"`
  Segments [1]SFNTLookupSegment `v8:"segments"`
}
