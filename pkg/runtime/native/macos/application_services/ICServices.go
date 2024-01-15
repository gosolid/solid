//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ICServices
*/

type ICServices struct {
  Count int16 `v8:"count"`
  Services [1]ICServiceEntry `v8:"services"`
}
