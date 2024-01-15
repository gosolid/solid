//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMScreeningChannelRec
*/

type CMScreeningChannelRec struct {
  Frequency int `v8:"frequency"`
  Angle int `v8:"angle"`
  SpotFunction uint `v8:"spotFunction"`
}
