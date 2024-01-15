//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMMultichannel5Color
*/

type CMMultichannel5Color struct {
  Components [5]byte `v8:"components"`
}
