//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_DB_RECORD_ATTRIBUTE_DATA_PTR
*/

type CSSMDBRECORDATTRIBUTEDATAPTR struct {
  DataRecordType uint `v8:"dataRecordType"`
  SemanticInformation uint `v8:"semanticInformation"`
  NumberOfAttributes uint `v8:"numberOfAttributes"`
  AttributeData *CSSMDBATTRIBUTEDATAPTR `v8:"attributeData"`
}
