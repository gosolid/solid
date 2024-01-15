//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMVideoCardGamma
*/

type CMVideoCardGamma struct {
  TagType uint `v8:"tagType"`
  U void `v8:"u"`
}
