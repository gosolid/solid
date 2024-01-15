//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.SpeechStatusInfo
*/

type SpeechStatusInfo struct {
  OutputBusy byte `v8:"outputBusy"`
  OutputPaused byte `v8:"outputPaused"`
  InputBytesLeft int64 `v8:"inputBytesLeft"`
  PhonemeCode int16 `v8:"phonemeCode"`
}
