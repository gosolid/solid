//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_QC_Statement
*/

type CEQCStatement struct {
  StatementId SecAsn1Oid `v8:"statementId"`
  SemanticsInfo CESemanticsInformation `v8:"semanticsInfo"`
  OtherInfo SecAsn1Oid `v8:"otherInfo"`
}
