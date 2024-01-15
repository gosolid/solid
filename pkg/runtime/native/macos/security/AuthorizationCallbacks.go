//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.AuthorizationCallbacks
*/

type AuthorizationCallbacks struct {
  Version uint `v8:"version"`
  SetResult func(...any) (any, error) `v8:"setResult"`
  RequestInterrupt func(...any) (any, error) `v8:"requestInterrupt"`
  DidDeactivate func(...any) (any, error) `v8:"didDeactivate"`
  GetContextValue func(...any) (any, error) `v8:"getContextValue"`
  SetContextValue func(...any) (any, error) `v8:"setContextValue"`
  GetHintValue func(...any) (any, error) `v8:"getHintValue"`
  SetHintValue func(...any) (any, error) `v8:"setHintValue"`
  GetArguments func(...any) (any, error) `v8:"getArguments"`
  GetSessionId func(...any) (any, error) `v8:"getSessionId"`
  GetImmutableHintValue func(...any) (any, error) `v8:"getImmutableHintValue"`
  GetLAContext func(...any) (any, error) `v8:"getLAContext"`
  GetTokenIdentities func(...any) (any, error) `v8:"getTokenIdentities"`
  GetTKTokenWatcher func(...any) (any, error) `v8:"getTKTokenWatcher"`
  RemoveHintValue func(...any) (any, error) `v8:"removeHintValue"`
  RemoveContextValue func(...any) (any, error) `v8:"removeContextValue"`
}
