//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509_REVOKED_CERT_ENTRY_PTR
*/

type CSSMX509REVOKEDCERTENTRYPTR struct {
  CertificateSerialNumber SecAsn1Oid `v8:"certificateSerialNumber"`
  RevocationDate CSSMX509TIMEPTR `v8:"revocationDate"`
  Extensions CSSMX509EXTENSIONSPTR `v8:"extensions"`
}
