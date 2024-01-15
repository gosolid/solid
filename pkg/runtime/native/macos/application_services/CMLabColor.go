//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMLabColor
*/

type CMLabColor struct {
  L uint16 `v8:"l"`
  A uint16 `v8:"a"`
  B uint16 `v8:"b"`
}
