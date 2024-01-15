//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMProfileLocation
*/

type CMProfileLocation struct {
  LocType int16 `v8:"locType"`
  U void `v8:"u"`
}
