//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMHandleLocation
*/

type CMHandleLocation struct {
  H **byte `v8:"h"`
}
