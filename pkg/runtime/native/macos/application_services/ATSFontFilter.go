//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSFontFilter
*/

type ATSFontFilter struct {
  Version uint `v8:"version"`
  FilterSelector int `v8:"filterSelector"`
  Filter void `v8:"filter"`
}
