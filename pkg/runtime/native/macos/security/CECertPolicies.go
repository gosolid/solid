//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_CertPolicies
*/

type CECertPolicies struct {
  NumPolicies uint `v8:"numPolicies"`
  Policies CEPolicyInformation `v8:"policies"`
}
