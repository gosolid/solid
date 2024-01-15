//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.SFNTLookupArrayHeader
*/

type SFNTLookupArrayHeader struct {
  LookupValues [1]uint16 `v8:"lookupValues"`
}
