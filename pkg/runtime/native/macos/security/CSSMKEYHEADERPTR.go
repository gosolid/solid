//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_KEYHEADER_PTR
*/

type CSSMKEYHEADERPTR struct {
  HeaderVersion uint `v8:"headerVersion"`
  CspId CSSMGUIDPTR `v8:"cspId"`
  BlobType uint `v8:"blobType"`
  Format uint `v8:"format"`
  AlgorithmId uint `v8:"algorithmId"`
  KeyClass uint `v8:"keyClass"`
  LogicalKeySizeInBits uint `v8:"logicalKeySizeInBits"`
  KeyAttr uint `v8:"keyAttr"`
  KeyUsage uint `v8:"keyUsage"`
  StartDate CSSMDATEPTR `v8:"startDate"`
  EndDate CSSMDATEPTR `v8:"endDate"`
  WrapAlgorithmId uint `v8:"wrapAlgorithmId"`
  WrapMode uint `v8:"wrapMode"`
  Reserved uint `v8:"reserved"`
}
