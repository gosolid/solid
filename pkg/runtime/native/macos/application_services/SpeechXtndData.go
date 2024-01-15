//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.SpeechXtndData
*/

type SpeechXtndData struct {
  SynthCreator uint `v8:"synthCreator"`
  SynthData [2]byte `v8:"synthData"`
}
