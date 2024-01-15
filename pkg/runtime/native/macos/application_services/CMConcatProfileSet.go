//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMConcatProfileSet
*/

type CMConcatProfileSet struct {
  KeyIndex uint16 `v8:"keyIndex"`
  Count uint16 `v8:"count"`
  ProfileSet [1]*OpaqueCMProfileRef `v8:"profileSet"`
}
