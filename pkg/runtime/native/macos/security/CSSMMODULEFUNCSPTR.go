//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_MODULE_FUNCS_PTR
*/

type CSSMMODULEFUNCSPTR struct {
  ServiceType uint `v8:"serviceType"`
  NumberOfServiceFuncs uint `v8:"numberOfServiceFuncs"`
  ServiceFuncs *func(...any) (any, error) `v8:"serviceFuncs"`
}
