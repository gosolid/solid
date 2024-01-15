//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_SELECTION_PREDICATE_PTR
*/

type CSSMSELECTIONPREDICATEPTR struct {
  DbOperator uint `v8:"dbOperator"`
  Attribute CSSMDBATTRIBUTEDATAPTR `v8:"attribute"`
}
