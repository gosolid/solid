//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.FMFontFamilyInstance
*/

type FMFontFamilyInstance struct {
  FontFamily int16 `v8:"fontFamily"`
  FontStyle int16 `v8:"fontStyle"`
}
