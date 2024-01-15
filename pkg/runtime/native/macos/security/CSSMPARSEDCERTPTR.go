//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_PARSED_CERT_PTR
*/

type CSSMPARSEDCERTPTR struct {
  CertType uint `v8:"certType"`
  ParsedCertFormat uint `v8:"parsedCertFormat"`
  ParsedCert void `v8:"parsedCert"`
}
