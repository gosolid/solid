//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSLayoutRecord
*/

type ATSLayoutRecord struct {
  GlyphID uint16 `v8:"glyphID"`
  Flags uint `v8:"flags"`
  OriginalOffset uint64 `v8:"originalOffset"`
  RealPos int `v8:"realPos"`
}
