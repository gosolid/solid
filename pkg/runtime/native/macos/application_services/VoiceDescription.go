//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.VoiceDescription
*/

type VoiceDescription struct {
  Length int `v8:"length"`
  Voice VoiceSpec `v8:"voice"`
  Version int `v8:"version"`
  Name [64]byte `v8:"name"`
  Comment [256]byte `v8:"comment"`
  Gender int16 `v8:"gender"`
  Age int16 `v8:"age"`
  Script int16 `v8:"script"`
  Language int16 `v8:"language"`
  Region int16 `v8:"region"`
  Reserved [4]int `v8:"reserved"`
}
