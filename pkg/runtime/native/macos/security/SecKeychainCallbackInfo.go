//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.SecKeychainCallbackInfo
*/

type SecKeychainCallbackInfo struct {
  Version uint `v8:"version"`
  Item *SecKeychainItem `v8:"item"`
  Keychain *SecKeychain `v8:"keychain"`
  Pid int `v8:"pid"`
}
