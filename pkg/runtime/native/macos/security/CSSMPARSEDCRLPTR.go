//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_PARSED_CRL_PTR
*/

type CSSMPARSEDCRLPTR struct {
  CrlType uint `v8:"crlType"`
  ParsedCrlFormat uint `v8:"parsedCrlFormat"`
  ParsedCrl void `v8:"parsedCrl"`
}
