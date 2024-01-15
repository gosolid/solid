//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSUGlyphInfo
*/

type ATSUGlyphInfo struct {
  GlyphID uint16 `v8:"glyphID"`
  Reserved uint16 `v8:"reserved"`
  LayoutFlags uint `v8:"layoutFlags"`
  CharIndex uint64 `v8:"charIndex"`
  Style *OpaqueATSUStyle `v8:"style"`
  DeltaY float32 `v8:"deltaY"`
  IdealX float32 `v8:"idealX"`
  ScreenX int16 `v8:"screenX"`
  CaretX int16 `v8:"caretX"`
}
