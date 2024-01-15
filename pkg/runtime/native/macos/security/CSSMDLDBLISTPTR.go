//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_DL_DB_LIST_PTR
*/

type CSSMDLDBLISTPTR struct {
  NumHandles uint `v8:"numHandles"`
  DLDBHandle *CSSMDLDBHANDLEPTR `v8:"dLDBHandle"`
}
