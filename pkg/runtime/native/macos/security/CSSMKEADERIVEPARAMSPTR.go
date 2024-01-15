//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_KEA_DERIVE_PARAMS_PTR
*/

type CSSMKEADERIVEPARAMSPTR struct {
  Rb SecAsn1Oid `v8:"rb"`
  Yb SecAsn1Oid `v8:"yb"`
}
