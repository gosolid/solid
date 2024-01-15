//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_CERT_PAIR_PTR
*/

type CSSMCERTPAIRPTR struct {
  EncodedCert CSSMENCODEDCERTPTR `v8:"encodedCert"`
  ParsedCert CSSMPARSEDCERTPTR `v8:"parsedCert"`
}
