//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.KernPair
*/

type KernPair struct {
  KernFirst byte `v8:"kernFirst"`
  KernSecond byte `v8:"kernSecond"`
  KernWidth int16 `v8:"kernWidth"`
}
