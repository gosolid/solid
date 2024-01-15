//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_APPLECSPDL_DB_IS_LOCKED_PARAMETERS_PTR
*/

type CSSMAPPLECSPDLDBISLOCKEDPARAMETERSPTR struct {
  IsLocked byte `v8:"isLocked"`
}
