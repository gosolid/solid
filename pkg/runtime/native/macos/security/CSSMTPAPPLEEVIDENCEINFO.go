//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_APPLE_EVIDENCE_INFO
*/

type CSSMTPAPPLEEVIDENCEINFO struct {
  StatusBits uint `v8:"statusBits"`
  NumStatusCodes uint `v8:"numStatusCodes"`
  StatusCodes int `v8:"statusCodes"`
  Index uint `v8:"index"`
  DlDbHandle CSSMDLDBHANDLEPTR `v8:"dlDbHandle"`
  UniqueRecord *CSSMDBUNIQUERECORDPTR `v8:"uniqueRecord"`
}
