//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_DB_SCHEMA_ATTRIBUTE_INFO_PTR
*/

type CSSMDBSCHEMAATTRIBUTEINFOPTR struct {
  AttributeId uint `v8:"attributeId"`
  AttributeName byte `v8:"attributeName"`
  AttributeNameID SecAsn1Oid `v8:"attributeNameID"`
  DataType uint `v8:"dataType"`
}
