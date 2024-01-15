//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_DB_RECORD_INDEX_INFO_PTR
*/

type CSSMDBRECORDINDEXINFOPTR struct {
  DataRecordType uint `v8:"dataRecordType"`
  NumberOfIndexes uint `v8:"numberOfIndexes"`
  IndexInfo *CSSMDBINDEXINFOPTR `v8:"indexInfo"`
}
