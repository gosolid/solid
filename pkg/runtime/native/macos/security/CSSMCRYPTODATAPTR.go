//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_CRYPTO_DATA_PTR
*/

type CSSMCRYPTODATAPTR struct {
  Param SecAsn1Oid `v8:"param"`
  Callback *func(...any) (any, error) `v8:"callback"`
  CallerCtx void `v8:"callerCtx"`
}
