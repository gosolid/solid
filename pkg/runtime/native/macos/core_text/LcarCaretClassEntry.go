//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.LcarCaretClassEntry
*/

type LcarCaretClassEntry struct {
  Count uint16 `v8:"count"`
  Partials [1]uint16 `v8:"partials"`
}
