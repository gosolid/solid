//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.FontInfo
*/

type FontInfo struct {
  Ascent int16 `v8:"ascent"`
  Descent int16 `v8:"descent"`
  WidMax int16 `v8:"widMax"`
  Leading int16 `v8:"leading"`
}
