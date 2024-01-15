//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_AUTHORIZATIONGROUP_PTR
*/

type CSSMAUTHORIZATIONGROUPPTR struct {
  NumberOfAuthTags uint `v8:"numberOfAuthTags"`
  AuthTags int `v8:"authTags"`
}
