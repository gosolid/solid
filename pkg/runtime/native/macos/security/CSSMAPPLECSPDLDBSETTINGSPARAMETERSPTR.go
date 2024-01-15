//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_APPLECSPDL_DB_SETTINGS_PARAMETERS_PTR
*/

type CSSMAPPLECSPDLDBSETTINGSPARAMETERSPTR struct {
  IdleTimeout uint `v8:"idleTimeout"`
  LockOnSleep byte `v8:"lockOnSleep"`
}
