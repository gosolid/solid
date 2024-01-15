//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509_REVOKED_CERT_LIST_PTR
*/

type CSSMX509REVOKEDCERTLISTPTR struct {
  NumberOfRevokedCertEntries uint `v8:"numberOfRevokedCertEntries"`
  RevokedCertEntry *CSSMX509REVOKEDCERTENTRYPTR `v8:"revokedCertEntry"`
}
