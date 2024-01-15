//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_CERTRECLAIM_OUTPUT_PTR
*/

type CSSMTPCERTRECLAIMOUTPUTPTR struct {
  ReclaimStatus uint `v8:"reclaimStatus"`
  ReclaimedCertGroup *CSSMCERTGROUPPTR `v8:"reclaimedCertGroup"`
  KeyCacheHandle uint64 `v8:"keyCacheHandle"`
}
