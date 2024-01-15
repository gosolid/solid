//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_SPI_DL_FUNCS_PTR
*/

type CSSMSPIDLFUNCSPTR struct {
  DbOpen func(...any) (any, error) `v8:"dbOpen"`
  DbClose func(...any) (any, error) `v8:"dbClose"`
  DbCreate func(...any) (any, error) `v8:"dbCreate"`
  DbDelete func(...any) (any, error) `v8:"dbDelete"`
  CreateRelation func(...any) (any, error) `v8:"createRelation"`
  DestroyRelation func(...any) (any, error) `v8:"destroyRelation"`
  Authenticate func(...any) (any, error) `v8:"authenticate"`
  GetDbAcl func(...any) (any, error) `v8:"getDbAcl"`
  ChangeDbAcl func(...any) (any, error) `v8:"changeDbAcl"`
  GetDbOwner func(...any) (any, error) `v8:"getDbOwner"`
  ChangeDbOwner func(...any) (any, error) `v8:"changeDbOwner"`
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
  PassThrough func(...any) (any, error) `v8:"passThrough"`
}
