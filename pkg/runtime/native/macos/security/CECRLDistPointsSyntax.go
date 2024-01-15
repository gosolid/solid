//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_CRLDistPointsSyntax
*/

type CECRLDistPointsSyntax struct {
  NumDistPoints uint `v8:"numDistPoints"`
  DistPoints CECRLDistributionPoint `v8:"distPoints"`
}
