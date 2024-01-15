//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_SPI_KR_FUNCS_PTR
*/

type CSSMSPIKRFUNCSPTR struct {
  RegistrationRequest func(...any) (any, error) `v8:"registrationRequest"`
  RegistrationRetrieve func(...any) (any, error) `v8:"registrationRetrieve"`
  GenerateRecoveryFields func(...any) (any, error) `v8:"generateRecoveryFields"`
  ProcessRecoveryFields func(...any) (any, error) `v8:"processRecoveryFields"`
  RecoveryRequest func(...any) (any, error) `v8:"recoveryRequest"`
  RecoveryRetrieve func(...any) (any, error) `v8:"recoveryRetrieve"`
  GetRecoveredObject func(...any) (any, error) `v8:"getRecoveredObject"`
  RecoveryRequestAbort func(...any) (any, error) `v8:"recoveryRequestAbort"`
  PassThrough func(...any) (any, error) `v8:"passThrough"`
}
