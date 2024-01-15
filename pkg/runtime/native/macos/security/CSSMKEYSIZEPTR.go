//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_KEY_SIZE_PTR
*/

type CSSMKEYSIZEPTR struct {
  LogicalKeySizeInBits uint `v8:"logicalKeySizeInBits"`
  EffectiveKeySizeInBits uint `v8:"effectiveKeySizeInBits"`
}
