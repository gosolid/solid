//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_SPI_CL_FUNCS_PTR
*/

type CSSMSPICLFUNCSPTR struct {
  CertCreateTemplate func(...any) (any, error) `v8:"certCreateTemplate"`
  CertGetAllTemplateFields func(...any) (any, error) `v8:"certGetAllTemplateFields"`
  CertSign func(...any) (any, error) `v8:"certSign"`
  CertVerify func(...any) (any, error) `v8:"certVerify"`
  CertVerifyWithKey func(...any) (any, error) `v8:"certVerifyWithKey"`
  CertGetFirstFieldValue func(...any) (any, error) `v8:"certGetFirstFieldValue"`
  CertGetNextFieldValue func(...any) (any, error) `v8:"certGetNextFieldValue"`
  CertAbortQuery func(...any) (any, error) `v8:"certAbortQuery"`
  CertGetKeyInfo func(...any) (any, error) `v8:"certGetKeyInfo"`
  CertGetAllFields func(...any) (any, error) `v8:"certGetAllFields"`
  FreeFields func(...any) (any, error) `v8:"freeFields"`
  FreeFieldValue func(...any) (any, error) `v8:"freeFieldValue"`
  CertCache func(...any) (any, error) `v8:"certCache"`
  CertGetFirstCachedFieldValue func(...any) (any, error) `v8:"certGetFirstCachedFieldValue"`
  CertGetNextCachedFieldValue func(...any) (any, error) `v8:"certGetNextCachedFieldValue"`
  CertAbortCache func(...any) (any, error) `v8:"certAbortCache"`
  CertGroupToSignedBundle func(...any) (any, error) `v8:"certGroupToSignedBundle"`
  CertGroupFromVerifiedBundle func(...any) (any, error) `v8:"certGroupFromVerifiedBundle"`
  CertDescribeFormat func(...any) (any, error) `v8:"certDescribeFormat"`
  CrlCreateTemplate func(...any) (any, error) `v8:"crlCreateTemplate"`
  CrlSetFields func(...any) (any, error) `v8:"crlSetFields"`
  CrlAddCert func(...any) (any, error) `v8:"crlAddCert"`
  CrlRemoveCert func(...any) (any, error) `v8:"crlRemoveCert"`
  CrlSign func(...any) (any, error) `v8:"crlSign"`
  CrlVerify func(...any) (any, error) `v8:"crlVerify"`
  CrlVerifyWithKey func(...any) (any, error) `v8:"crlVerifyWithKey"`
  IsCertInCrl func(...any) (any, error) `v8:"isCertInCrl"`
  CrlGetFirstFieldValue func(...any) (any, error) `v8:"crlGetFirstFieldValue"`
  CrlGetNextFieldValue func(...any) (any, error) `v8:"crlGetNextFieldValue"`
  CrlAbortQuery func(...any) (any, error) `v8:"crlAbortQuery"`
  CrlGetAllFields func(...any) (any, error) `v8:"crlGetAllFields"`
  CrlCache func(...any) (any, error) `v8:"crlCache"`
  IsCertInCachedCrl func(...any) (any, error) `v8:"isCertInCachedCrl"`
  CrlGetFirstCachedFieldValue func(...any) (any, error) `v8:"crlGetFirstCachedFieldValue"`
  CrlGetNextCachedFieldValue func(...any) (any, error) `v8:"crlGetNextCachedFieldValue"`
  CrlGetAllCachedRecordFields func(...any) (any, error) `v8:"crlGetAllCachedRecordFields"`
  CrlAbortCache func(...any) (any, error) `v8:"crlAbortCache"`
  CrlDescribeFormat func(...any) (any, error) `v8:"crlDescribeFormat"`
  PassThrough func(...any) (any, error) `v8:"passThrough"`
}
