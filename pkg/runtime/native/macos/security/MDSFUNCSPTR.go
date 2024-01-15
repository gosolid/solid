//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.MDS_FUNCS_PTR
*/

type MDSFUNCSPTR struct {
  DbOpen func(...any) (any, error) `v8:"dbOpen"`
  DbClose func(...any) (any, error) `v8:"dbClose"`
  GetDbNames func(...any) (any, error) `v8:"getDbNames"`
  GetDbNameFromHandle func(...any) (any, error) `v8:"getDbNameFromHandle"`
  FreeNameList func(...any) (any, error) `v8:"freeNameList"`
  DataInsert func(...any) (any, error) `v8:"dataInsert"`
  DataDelete func(...any) (any, error) `v8:"dataDelete"`
  DataModify func(...any) (any, error) `v8:"dataModify"`
  DataGetFirst func(...any) (any, error) `v8:"dataGetFirst"`
  DataGetNext func(...any) (any, error) `v8:"dataGetNext"`
  DataAbortQuery func(...any) (any, error) `v8:"dataAbortQuery"`
  DataGetFromUniqueRecordId func(...any) (any, error) `v8:"dataGetFromUniqueRecordId"`
  FreeUniqueRecord func(...any) (any, error) `v8:"freeUniqueRecord"`
  CreateRelation func(...any) (any, error) `v8:"createRelation"`
  DestroyRelation func(...any) (any, error) `v8:"destroyRelation"`
}
