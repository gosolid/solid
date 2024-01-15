//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_DL_PKCS11_ATTRIBUTE_PTR
*/

type CSSMDLPKCS11ATTRIBUTEPTR struct {
  DeviceAccessFlags uint `v8:"deviceAccessFlags"`
}
