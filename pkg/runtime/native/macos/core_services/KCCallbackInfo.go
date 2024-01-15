//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/security"
)

/*
struct CoreServices.KCCallbackInfo
*/

type KCCallbackInfo struct {
  Version uint `v8:"version"`
  Item *security.SecKeychainItem `v8:"item"`
  ProcessID [2]int `v8:"processID"`
  Event [4]int `v8:"event"`
  Keychain *security.SecKeychain `v8:"keychain"`
}
