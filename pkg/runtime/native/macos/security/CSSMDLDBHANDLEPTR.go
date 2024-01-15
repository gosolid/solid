//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_DL_DB_HANDLE_PTR
*/

type CSSMDLDBHANDLEPTR struct {
  DLHandle int64 `v8:"dLHandle"`
  DBHandle int64 `v8:"dBHandle"`
}
