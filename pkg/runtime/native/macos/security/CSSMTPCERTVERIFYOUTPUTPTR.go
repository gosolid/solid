//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_CERTVERIFY_OUTPUT_PTR
*/

type CSSMTPCERTVERIFYOUTPUTPTR struct {
  VerifyStatus uint `v8:"verifyStatus"`
  NumberOfEvidence uint `v8:"numberOfEvidence"`
  Evidence *CSSMEVIDENCEPTR `v8:"evidence"`
}
