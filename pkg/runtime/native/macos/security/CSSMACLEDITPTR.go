//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_ACL_EDIT_PTR
*/

type CSSMACLEDITPTR struct {
  EditMode uint `v8:"editMode"`
  OldEntryHandle int64 `v8:"oldEntryHandle"`
  NewEntry CSSMACLENTRYINPUTPTR `v8:"newEntry"`
}
