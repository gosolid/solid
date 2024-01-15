//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.SecKeychainAttributeList
*/

type SecKeychainAttributeList struct {
  Count uint `v8:"count"`
  Attr SecKeychainAttribute `v8:"attr"`
}
