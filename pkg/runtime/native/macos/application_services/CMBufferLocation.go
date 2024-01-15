//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMBufferLocation
*/

type CMBufferLocation struct {
  Buffer void `v8:"buffer"`
  Size uint `v8:"size"`
}
