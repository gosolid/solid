//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMCMYColor
*/

type CMCMYColor struct {
  Cyan uint16 `v8:"cyan"`
  Magenta uint16 `v8:"magenta"`
  Yellow uint16 `v8:"yellow"`
}
