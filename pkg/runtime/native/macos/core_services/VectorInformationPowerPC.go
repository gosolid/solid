//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.VectorInformationPowerPC
*/

type VectorInformationPowerPC struct {
  Registers void `v8:"registers"`
  VSCR void `v8:"vSCR"`
  VRsave uint `v8:"vRsave"`
}
