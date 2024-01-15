//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.PhonemeDescriptor
*/

type PhonemeDescriptor struct {
  PhonemeCount int16 `v8:"phonemeCount"`
  ThePhonemes [1]PhonemeInfo `v8:"thePhonemes"`
}
