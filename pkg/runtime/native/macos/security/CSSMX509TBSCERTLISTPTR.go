//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509_TBS_CERTLIST_PTR
*/

type CSSMX509TBSCERTLISTPTR struct {
  Version SecAsn1Oid `v8:"version"`
  Signature SecAsn1AlgId `v8:"signature"`
  Issuer CSSMX509NAMEPTR `v8:"issuer"`
  ThisUpdate CSSMX509TIMEPTR `v8:"thisUpdate"`
  NextUpdate CSSMX509TIMEPTR `v8:"nextUpdate"`
  RevokedCertificates *CSSMX509REVOKEDCERTLISTPTR `v8:"revokedCertificates"`
  Extensions CSSMX509EXTENSIONSPTR `v8:"extensions"`
}
