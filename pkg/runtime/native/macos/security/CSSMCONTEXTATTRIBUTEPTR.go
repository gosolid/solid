//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_CONTEXT_ATTRIBUTE_PTR
*/

type CSSMCONTEXTATTRIBUTEPTR struct {
  AttributeType uint `v8:"attributeType"`
  AttributeLength uint `v8:"attributeLength"`
  Attribute void `v8:"attribute"`
}
