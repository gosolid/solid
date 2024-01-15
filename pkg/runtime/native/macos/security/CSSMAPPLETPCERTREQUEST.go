//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_APPLE_TP_CERT_REQUEST
*/

type CSSMAPPLETPCERTREQUEST struct {
  CspHand int64 `v8:"cspHand"`
  ClHand int64 `v8:"clHand"`
  SerialNumber uint `v8:"serialNumber"`
  NumSubjectNames uint `v8:"numSubjectNames"`
  SubjectNames CSSMAPPLETPNAMEOID `v8:"subjectNames"`
  NumIssuerNames uint `v8:"numIssuerNames"`
  IssuerNames CSSMAPPLETPNAMEOID `v8:"issuerNames"`
  IssuerNameX509 *CSSMX509NAMEPTR `v8:"issuerNameX509"`
  CertPublicKey CSSMKEYPTR `v8:"certPublicKey"`
  IssuerPrivateKey CSSMKEYPTR `v8:"issuerPrivateKey"`
  SignatureAlg uint `v8:"signatureAlg"`
  SignatureOid SecAsn1Oid `v8:"signatureOid"`
  NotBefore uint `v8:"notBefore"`
  NotAfter uint `v8:"notAfter"`
  NumExtensions uint `v8:"numExtensions"`
  Extensions CEDataAndType `v8:"extensions"`
  ChallengeString byte `v8:"challengeString"`
}
