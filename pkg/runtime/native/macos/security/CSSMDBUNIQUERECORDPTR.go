//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_DB_UNIQUE_RECORD_PTR
*/

type CSSMDBUNIQUERECORDPTR struct {
  RecordLocator CSSMDBINDEXINFOPTR `v8:"recordLocator"`
  RecordIdentifier SecAsn1Oid `v8:"recordIdentifier"`
}
