//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_CERTCHANGE_OUTPUT_PTR
*/

type CSSMTPCERTCHANGEOUTPUTPTR struct {
  ActionStatus uint `v8:"actionStatus"`
  RevokeInfo CSSMFIELDPTR `v8:"revokeInfo"`
}
