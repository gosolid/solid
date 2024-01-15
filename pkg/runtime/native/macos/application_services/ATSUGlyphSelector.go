//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSUGlyphSelector
*/

type ATSUGlyphSelector struct {
  Collection uint16 `v8:"collection"`
  GlyphID uint16 `v8:"glyphID"`
}
