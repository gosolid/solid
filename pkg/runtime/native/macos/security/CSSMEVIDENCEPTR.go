//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_EVIDENCE_PTR
*/

type CSSMEVIDENCEPTR struct {
  EvidenceForm uint `v8:"evidenceForm"`
  Evidence void `v8:"evidence"`
}
