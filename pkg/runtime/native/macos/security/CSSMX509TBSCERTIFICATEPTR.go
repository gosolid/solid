//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509_TBS_CERTIFICATE_PTR
*/

type CSSMX509TBSCERTIFICATEPTR struct {
  Version SecAsn1Oid `v8:"version"`
  SerialNumber SecAsn1Oid `v8:"serialNumber"`
  Signature SecAsn1AlgId `v8:"signature"`
  Issuer CSSMX509NAMEPTR `v8:"issuer"`
  Validity CSSMX509VALIDITYPTR `v8:"validity"`
  Subject CSSMX509NAMEPTR `v8:"subject"`
  SubjectPublicKeyInfo SecAsn1PubKeyInfo `v8:"subjectPublicKeyInfo"`
  IssuerUniqueIdentifier SecAsn1Oid `v8:"issuerUniqueIdentifier"`
  SubjectUniqueIdentifier SecAsn1Oid `v8:"subjectUniqueIdentifier"`
  Extensions CSSMX509EXTENSIONSPTR `v8:"extensions"`
}
