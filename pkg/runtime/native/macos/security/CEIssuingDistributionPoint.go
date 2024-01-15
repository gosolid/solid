//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_IssuingDistributionPoint
*/

type CEIssuingDistributionPoint struct {
  DistPointName CEDistributionPointName `v8:"distPointName"`
  OnlyUserCertsPresent int `v8:"onlyUserCertsPresent"`
  OnlyUserCerts int `v8:"onlyUserCerts"`
  OnlyCACertsPresent int `v8:"onlyCACertsPresent"`
  OnlyCACerts int `v8:"onlyCACerts"`
  OnlySomeReasonsPresent int `v8:"onlySomeReasonsPresent"`
  OnlySomeReasons byte `v8:"onlySomeReasons"`
  IndirectCrlPresent int `v8:"indirectCrlPresent"`
  IndirectCrl int `v8:"indirectCrl"`
}
