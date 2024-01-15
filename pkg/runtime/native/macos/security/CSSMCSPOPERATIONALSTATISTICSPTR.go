//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_CSP_OPERATIONAL_STATISTICS_PTR
*/

type CSSMCSPOPERATIONALSTATISTICSPTR struct {
  UserAuthenticated int `v8:"userAuthenticated"`
  DeviceFlags uint `v8:"deviceFlags"`
  TokenMaxSessionCount uint `v8:"tokenMaxSessionCount"`
  TokenOpenedSessionCount uint `v8:"tokenOpenedSessionCount"`
  TokenMaxRWSessionCount uint `v8:"tokenMaxRWSessionCount"`
  TokenOpenedRWSessionCount uint `v8:"tokenOpenedRWSessionCount"`
  TokenTotalPublicMem uint `v8:"tokenTotalPublicMem"`
  TokenFreePublicMem uint `v8:"tokenFreePublicMem"`
  TokenTotalPrivateMem uint `v8:"tokenTotalPrivateMem"`
  TokenFreePrivateMem uint `v8:"tokenFreePrivateMem"`
}
