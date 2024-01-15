//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct ApplicationServices.SpeechVersionInfo
*/

type SpeechVersionInfo struct {
  SynthType uint `v8:"synthType"`
  SynthSubType uint `v8:"synthSubType"`
  SynthManufacturer uint `v8:"synthManufacturer"`
  SynthFlags int `v8:"synthFlags"`
  SynthVersion objc.NumVersion `v8:"synthVersion"`
}
