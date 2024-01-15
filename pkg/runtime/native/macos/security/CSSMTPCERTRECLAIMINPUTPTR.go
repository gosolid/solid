//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_CERTRECLAIM_INPUT_PTR
*/

type CSSMTPCERTRECLAIMINPUTPTR struct {
  CLHandle int64 `v8:"cLHandle"`
  NumberOfSelectionFields uint `v8:"numberOfSelectionFields"`
  SelectionFields *CSSMFIELDPTR `v8:"selectionFields"`
  UserCredentials *CSSMACCESSCREDENTIALSPTR `v8:"userCredentials"`
}
