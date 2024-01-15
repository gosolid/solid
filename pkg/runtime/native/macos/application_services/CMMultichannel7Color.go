//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMMultichannel7Color
*/

type CMMultichannel7Color struct {
  Components [7]byte `v8:"components"`
}
