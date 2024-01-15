//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_APPLECSPDL_DB_CHANGE_PASSWORD_PARAMETERS_PTR
*/

type CSSMAPPLECSPDLDBCHANGEPASSWORDPARAMETERSPTR struct {
  AccessCredentials CSSMACCESSCREDENTIALSPTR `v8:"accessCredentials"`
}
