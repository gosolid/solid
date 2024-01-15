//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.SpeechChannelRecord
*/

type SpeechChannelRecord struct {
  Data [1]int64 `v8:"data"`
}
