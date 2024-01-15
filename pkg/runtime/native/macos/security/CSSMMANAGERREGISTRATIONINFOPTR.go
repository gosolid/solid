//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_MANAGER_REGISTRATION_INFO_PTR
*/

type CSSMMANAGERREGISTRATIONINFOPTR struct {
  Initialize func(...any) (any, error) `v8:"initialize"`
  Terminate func(...any) (any, error) `v8:"terminate"`
  RegisterDispatchTable func(...any) (any, error) `v8:"registerDispatchTable"`
  DeregisterDispatchTable func(...any) (any, error) `v8:"deregisterDispatchTable"`
  EventNotifyManager func(...any) (any, error) `v8:"eventNotifyManager"`
  RefreshFunctionTable func(...any) (any, error) `v8:"refreshFunctionTable"`
}
