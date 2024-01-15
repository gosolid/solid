//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSUStyleRunInfo
*/

type ATSUStyleRunInfo struct {
  RunLength uint `v8:"runLength"`
  StyleObjectIndex uint `v8:"styleObjectIndex"`
}
