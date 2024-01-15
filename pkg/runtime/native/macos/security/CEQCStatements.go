//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_QC_Statements
*/

type CEQCStatements struct {
  NumQCStatements uint `v8:"numQCStatements"`
  QcStatements CEQCStatement `v8:"qcStatements"`
}
