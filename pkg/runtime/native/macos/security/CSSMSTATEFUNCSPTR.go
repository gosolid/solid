//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_STATE_FUNCS_PTR
*/

type CSSMSTATEFUNCSPTR struct {
  CssmGetAttachFunctions func(...any) (any, error) `v8:"cssmGetAttachFunctions"`
  CssmReleaseAttachFunctions func(...any) (any, error) `v8:"cssmReleaseAttachFunctions"`
  CssmGetAppMemoryFunctions func(...any) (any, error) `v8:"cssmGetAppMemoryFunctions"`
  CssmIsFuncCallValid func(...any) (any, error) `v8:"cssmIsFuncCallValid"`
  CssmDeregisterManagerServices func(...any) (any, error) `v8:"cssmDeregisterManagerServices"`
  CssmDeliverModuleManagerEvent func(...any) (any, error) `v8:"cssmDeliverModuleManagerEvent"`
}
