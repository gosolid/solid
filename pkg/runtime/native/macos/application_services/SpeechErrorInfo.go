//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.SpeechErrorInfo
*/

type SpeechErrorInfo struct {
  Count int16 `v8:"count"`
  Oldest int16 `v8:"oldest"`
  OldPos int64 `v8:"oldPos"`
  Newest int16 `v8:"newest"`
  NewPos int64 `v8:"newPos"`
}
