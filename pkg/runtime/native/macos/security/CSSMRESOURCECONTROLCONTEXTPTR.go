//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_RESOURCE_CONTROL_CONTEXT_PTR
*/

type CSSMRESOURCECONTROLCONTEXTPTR struct {
  AccessCred *CSSMACCESSCREDENTIALSPTR `v8:"accessCred"`
  InitialAclEntry CSSMACLENTRYINPUTPTR `v8:"initialAclEntry"`
}
