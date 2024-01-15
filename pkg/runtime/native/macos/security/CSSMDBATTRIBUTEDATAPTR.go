//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_DB_ATTRIBUTE_DATA_PTR
*/

type CSSMDBATTRIBUTEDATAPTR struct {
  Info CSSMDBATTRIBUTEINFOPTR `v8:"info"`
  NumberOfValues uint `v8:"numberOfValues"`
  Value *SecAsn1Oid `v8:"value"`
}
