//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_SPI_TP_FUNCS_PTR
*/

type CSSMSPITPFUNCSPTR struct {
  SubmitCredRequest func(...any) (any, error) `v8:"submitCredRequest"`
  RetrieveCredResult func(...any) (any, error) `v8:"retrieveCredResult"`
  ConfirmCredResult func(...any) (any, error) `v8:"confirmCredResult"`
  ReceiveConfirmation func(...any) (any, error) `v8:"receiveConfirmation"`
  CertReclaimKey func(...any) (any, error) `v8:"certReclaimKey"`
  CertReclaimAbort func(...any) (any, error) `v8:"certReclaimAbort"`
  FormRequest func(...any) (any, error) `v8:"formRequest"`
  FormSubmit func(...any) (any, error) `v8:"formSubmit"`
  CertGroupVerify func(...any) (any, error) `v8:"certGroupVerify"`
  CertCreateTemplate func(...any) (any, error) `v8:"certCreateTemplate"`
  CertGetAllTemplateFields func(...any) (any, error) `v8:"certGetAllTemplateFields"`
  CertSign func(...any) (any, error) `v8:"certSign"`
  CrlVerify func(...any) (any, error) `v8:"crlVerify"`
  CrlCreateTemplate func(...any) (any, error) `v8:"crlCreateTemplate"`
  CertRevoke func(...any) (any, error) `v8:"certRevoke"`
  CertRemoveFromCrlTemplate func(...any) (any, error) `v8:"certRemoveFromCrlTemplate"`
  CrlSign func(...any) (any, error) `v8:"crlSign"`
  ApplyCrlToDb func(...any) (any, error) `v8:"applyCrlToDb"`
  CertGroupConstruct func(...any) (any, error) `v8:"certGroupConstruct"`
  CertGroupPrune func(...any) (any, error) `v8:"certGroupPrune"`
  CertGroupToTupleGroup func(...any) (any, error) `v8:"certGroupToTupleGroup"`
  TupleGroupToCertGroup func(...any) (any, error) `v8:"tupleGroupToCertGroup"`
  PassThrough func(...any) (any, error) `v8:"passThrough"`
}
