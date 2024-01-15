//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSUGlyphInfoArray
*/

type ATSUGlyphInfoArray struct {
  Layout *OpaqueATSUTextLayout `v8:"layout"`
  NumGlyphs uint64 `v8:"numGlyphs"`
  Glyphs [1]ATSUGlyphInfo `v8:"glyphs"`
}
