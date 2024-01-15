//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_AUTHORITY_ID_PTR
*/

type CSSMTPAUTHORITYIDPTR struct {
  AuthorityCert SecAsn1Oid `v8:"authorityCert"`
  AuthorityLocation *CSSMNETADDRESSPTR `v8:"authorityLocation"`
}
