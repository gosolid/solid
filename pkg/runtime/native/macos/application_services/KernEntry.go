//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.KernEntry
*/

type KernEntry struct {
  KernStyle int16 `v8:"kernStyle"`
  KernLength int16 `v8:"kernLength"`
}
