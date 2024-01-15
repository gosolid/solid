//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ICCharTable
*/

type ICCharTable struct {
  NetToMac [256]byte `v8:"netToMac"`
  MacToNet [256]byte `v8:"macToNet"`
}
