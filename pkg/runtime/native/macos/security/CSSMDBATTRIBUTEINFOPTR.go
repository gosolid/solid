//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_DB_ATTRIBUTE_INFO_PTR
*/

type CSSMDBATTRIBUTEINFOPTR struct {
  AttributeNameFormat uint `v8:"attributeNameFormat"`
  Label void `v8:"label"`
  AttributeFormat uint `v8:"attributeFormat"`
}
