//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.SecKeychainAttributeInfo
*/

type SecKeychainAttributeInfo struct {
  Count uint `v8:"count"`
  Tag uint `v8:"tag"`
  Format uint `v8:"format"`
}
