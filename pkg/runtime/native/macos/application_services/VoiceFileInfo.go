//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_services"
)

/*
struct ApplicationServices.VoiceFileInfo
*/

type VoiceFileInfo struct {
  FileSpec core_services.FSSpec `v8:"fileSpec"`
  ResID int16 `v8:"resID"`
}
