//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMMultichannel6Color
*/

type CMMultichannel6Color struct {
  Components [6]byte `v8:"components"`
}
