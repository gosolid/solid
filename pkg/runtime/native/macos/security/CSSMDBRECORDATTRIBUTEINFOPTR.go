//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_DB_RECORD_ATTRIBUTE_INFO_PTR
*/

type CSSMDBRECORDATTRIBUTEINFOPTR struct {
  DataRecordType uint `v8:"dataRecordType"`
  NumberOfAttributes uint `v8:"numberOfAttributes"`
  AttributeInfo *CSSMDBATTRIBUTEINFOPTR `v8:"attributeInfo"`
}
