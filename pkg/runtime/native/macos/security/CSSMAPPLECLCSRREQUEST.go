//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_APPLE_CL_CSR_REQUEST
*/

type CSSMAPPLECLCSRREQUEST struct {
  SubjectNameX509 *CSSMX509NAMEPTR `v8:"subjectNameX509"`
  SignatureAlg uint `v8:"signatureAlg"`
  SignatureOid SecAsn1Oid `v8:"signatureOid"`
  CspHand int64 `v8:"cspHand"`
  SubjectPublicKey CSSMKEYPTR `v8:"subjectPublicKey"`
  SubjectPrivateKey CSSMKEYPTR `v8:"subjectPrivateKey"`
  ChallengeString byte `v8:"challengeString"`
}
