//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMMultichannel8Color
*/

type CMMultichannel8Color struct {
  Components [8]byte `v8:"components"`
}
