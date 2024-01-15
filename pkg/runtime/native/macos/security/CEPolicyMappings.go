//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_PolicyMappings
*/

type CEPolicyMappings struct {
  NumPolicyMappings uint `v8:"numPolicyMappings"`
  PolicyMappings CEPolicyMapping `v8:"policyMappings"`
}
