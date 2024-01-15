//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_LIST_PTR
*/

type CSSMLISTPTR struct {
  ListType uint `v8:"listType"`
  Head *CSSMLISTELEMENT `v8:"head"`
  Tail *CSSMLISTELEMENT `v8:"tail"`
}
