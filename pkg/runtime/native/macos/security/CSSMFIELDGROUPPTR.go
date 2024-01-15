//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_FIELDGROUP_PTR
*/

type CSSMFIELDGROUPPTR struct {
  NumberOfFields int `v8:"numberOfFields"`
  Fields *CSSMFIELDPTR `v8:"fields"`
}
