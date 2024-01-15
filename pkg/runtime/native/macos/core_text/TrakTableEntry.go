//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.TrakTableEntry
*/

type TrakTableEntry struct {
  Track int `v8:"track"`
  NameTableIndex uint16 `v8:"nameTableIndex"`
  SizesOffset uint16 `v8:"sizesOffset"`
}
