//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_ACL_OWNER_PROTOTYPE_PTR
*/

type CSSMACLOWNERPROTOTYPEPTR struct {
  TypedSubject CSSMLISTPTR `v8:"typedSubject"`
  Delegate int `v8:"delegate"`
}
