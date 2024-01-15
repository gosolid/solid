//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.PMLanguageInfo
*/

type PMLanguageInfo struct {
  Level [33]byte `v8:"level"`
  Version [33]byte `v8:"version"`
  Release [33]byte `v8:"release"`
}
