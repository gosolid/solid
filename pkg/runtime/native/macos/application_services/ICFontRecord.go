//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ICFontRecord
*/

type ICFontRecord struct {
  Size int16 `v8:"size"`
  Face byte `v8:"face"`
  Pad byte `v8:"pad"`
  Font [256]byte `v8:"font"`
}
