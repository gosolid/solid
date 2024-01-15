//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_KR_POLICY_LIST_ITEM_PTR
*/

type CSSMKRPOLICYLISTITEMPTR struct {
  Next KrPolicyListItem `v8:"next"`
  AlgorithmId uint `v8:"algorithmId"`
  Mode uint `v8:"mode"`
  MaxKeyLength uint `v8:"maxKeyLength"`
  MaxRounds uint `v8:"maxRounds"`
  WorkFactor byte `v8:"workFactor"`
  PolicyFlags uint `v8:"policyFlags"`
  AlgClass uint `v8:"algClass"`
}
