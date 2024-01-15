//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.SecKeychainSettings
*/

type SecKeychainSettings struct {
  Version uint `v8:"version"`
  LockOnSleep byte `v8:"lockOnSleep"`
  UseLockInterval byte `v8:"useLockInterval"`
  LockInterval uint `v8:"lockInterval"`
}
