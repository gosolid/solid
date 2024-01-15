//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_ACCESS_CREDENTIALS_PTR
*/

type CSSMACCESSCREDENTIALSPTR struct {
  EntryTag [68]byte `v8:"entryTag"`
  BaseCerts CSSMBASECERTSPTR `v8:"baseCerts"`
  Samples CSSMSAMPLEGROUPPTR `v8:"samples"`
  Callback *func(...any) (any, error) `v8:"callback"`
  CallerCtx void `v8:"callerCtx"`
}
