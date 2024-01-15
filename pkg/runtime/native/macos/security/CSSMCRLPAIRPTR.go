//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_CRL_PAIR_PTR
*/

type CSSMCRLPAIRPTR struct {
  EncodedCrl CSSMENCODEDCRLPTR `v8:"encodedCrl"`
  ParsedCrl CSSMPARSEDCRLPTR `v8:"parsedCrl"`
}
