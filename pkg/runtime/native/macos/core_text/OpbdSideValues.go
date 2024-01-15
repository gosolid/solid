//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.OpbdSideValues
*/

type OpbdSideValues struct {
  LeftSideShift int16 `v8:"leftSideShift"`
  TopSideShift int16 `v8:"topSideShift"`
  RightSideShift int16 `v8:"rightSideShift"`
  BottomSideShift int16 `v8:"bottomSideShift"`
}
