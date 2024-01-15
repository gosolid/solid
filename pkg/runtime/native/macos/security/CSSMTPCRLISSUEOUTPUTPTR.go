//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_CRLISSUE_OUTPUT_PTR
*/

type CSSMTPCRLISSUEOUTPUTPTR struct {
  IssueStatus uint `v8:"issueStatus"`
  Crl *CSSMENCODEDCRLPTR `v8:"crl"`
  CrlNextTime *byte `v8:"crlNextTime"`
}
