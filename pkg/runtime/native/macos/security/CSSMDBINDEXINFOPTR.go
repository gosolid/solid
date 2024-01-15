//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_DB_INDEX_INFO_PTR
*/

type CSSMDBINDEXINFOPTR struct {
  IndexType uint `v8:"indexType"`
  IndexedDataLocation uint `v8:"indexedDataLocation"`
  Info CSSMDBATTRIBUTEINFOPTR `v8:"info"`
}
