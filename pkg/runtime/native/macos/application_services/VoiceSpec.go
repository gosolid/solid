//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.VoiceSpec
*/

type VoiceSpec struct {
  Creator uint `v8:"creator"`
  Id uint `v8:"id"`
}
