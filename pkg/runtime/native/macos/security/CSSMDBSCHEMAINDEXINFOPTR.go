//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_DB_SCHEMA_INDEX_INFO_PTR
*/

type CSSMDBSCHEMAINDEXINFOPTR struct {
  AttributeId uint `v8:"attributeId"`
  IndexId uint `v8:"indexId"`
  IndexType uint `v8:"indexType"`
  IndexedDataLocation uint `v8:"indexedDataLocation"`
}
