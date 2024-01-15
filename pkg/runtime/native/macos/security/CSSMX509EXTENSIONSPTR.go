//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509_EXTENSIONS_PTR
*/

type CSSMX509EXTENSIONSPTR struct {
  NumberOfExtensions uint `v8:"numberOfExtensions"`
  Extensions *CSSMX509EXTENSIONPTR `v8:"extensions"`
}
