//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_VERIFY_CONTEXT_RESULT_PTR
*/

type CSSMTPVERIFYCONTEXTRESULTPTR struct {
  NumberOfEvidences uint `v8:"numberOfEvidences"`
  Evidence *CSSMEVIDENCEPTR `v8:"evidence"`
}
