//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_CRLDistributionPoint
*/

type CECRLDistributionPoint struct {
  DistPointName CEDistributionPointName `v8:"distPointName"`
  ReasonsPresent int `v8:"reasonsPresent"`
  Reasons byte `v8:"reasons"`
  CrlIssuer CEGeneralNames `v8:"crlIssuer"`
}
