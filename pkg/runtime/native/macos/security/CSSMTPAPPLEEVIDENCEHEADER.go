//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_APPLE_EVIDENCE_HEADER
*/

type CSSMTPAPPLEEVIDENCEHEADER struct {
  Version uint `v8:"version"`
}
