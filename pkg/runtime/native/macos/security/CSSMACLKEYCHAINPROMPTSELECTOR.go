//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_ACL_KEYCHAIN_PROMPT_SELECTOR
*/

type CSSMACLKEYCHAINPROMPTSELECTOR struct {
  Version uint16 `v8:"version"`
  Flags uint16 `v8:"flags"`
}
