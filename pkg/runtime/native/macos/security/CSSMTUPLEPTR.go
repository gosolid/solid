//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TUPLE_PTR
*/

type CSSMTUPLEPTR struct {
  Issuer CSSMLISTPTR `v8:"issuer"`
  Subject CSSMLISTPTR `v8:"subject"`
  Delegate int `v8:"delegate"`
  AuthorizationTag CSSMLISTPTR `v8:"authorizationTag"`
  ValidityPeriod CSSMLISTPTR `v8:"validityPeriod"`
}
