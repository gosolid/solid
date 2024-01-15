//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMNamedColor
*/

type CMNamedColor struct {
  NamedColorIndex uint `v8:"namedColorIndex"`
}
