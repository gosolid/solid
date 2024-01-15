//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.AnchorPointTable
*/

type AnchorPointTable struct {
  NPoints uint `v8:"nPoints"`
  Points [1]AnchorPoint `v8:"points"`
}
