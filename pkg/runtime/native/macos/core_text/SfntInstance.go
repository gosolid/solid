//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.sfntInstance
*/

type SfntInstance struct {
  NameID int16 `v8:"nameID"`
  Flags int16 `v8:"flags"`
  Coord [1]int `v8:"coord"`
}
