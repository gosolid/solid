//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_PolicyConstraints
*/

type CEPolicyConstraints struct {
  RequireExplicitPolicyPresent int `v8:"requireExplicitPolicyPresent"`
  RequireExplicitPolicy uint `v8:"requireExplicitPolicy"`
  InhibitPolicyMappingPresent int `v8:"inhibitPolicyMappingPresent"`
  InhibitPolicyMapping uint `v8:"inhibitPolicyMapping"`
}
