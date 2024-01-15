//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.MortSwashSubtable
*/

type MortSwashSubtable struct {
  Lookup SFNTLookupTable `v8:"lookup"`
}
