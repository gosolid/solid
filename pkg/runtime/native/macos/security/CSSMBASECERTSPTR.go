//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_BASE_CERTS_PTR
*/

type CSSMBASECERTSPTR struct {
  TPHandle int64 `v8:"tPHandle"`
  CLHandle int64 `v8:"cLHandle"`
  Certs CSSMCERTGROUPPTR `v8:"certs"`
}
