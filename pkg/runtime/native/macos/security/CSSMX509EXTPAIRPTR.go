//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509EXT_PAIR_PTR
*/

type CSSMX509EXTPAIRPTR struct {
  TagAndValue CSSMX509EXTTAGandVALUEPTR `v8:"tagAndValue"`
  ParsedValue void `v8:"parsedValue"`
}
