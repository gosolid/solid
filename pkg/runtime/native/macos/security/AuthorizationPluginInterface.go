//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.AuthorizationPluginInterface
*/

type AuthorizationPluginInterface struct {
  Version uint `v8:"version"`
  PluginDestroy func(...any) (any, error) `v8:"pluginDestroy"`
  MechanismCreate func(...any) (any, error) `v8:"mechanismCreate"`
  MechanismInvoke func(...any) (any, error) `v8:"mechanismInvoke"`
  MechanismDeactivate func(...any) (any, error) `v8:"mechanismDeactivate"`
  MechanismDestroy func(...any) (any, error) `v8:"mechanismDestroy"`
}
