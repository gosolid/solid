//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509_VALIDITY_PTR
*/

type CSSMX509VALIDITYPTR struct {
  NotBefore CSSMX509TIMEPTR `v8:"notBefore"`
  NotAfter CSSMX509TIMEPTR `v8:"notAfter"`
}
