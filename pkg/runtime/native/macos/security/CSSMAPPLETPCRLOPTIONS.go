//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_APPLE_TP_CRL_OPTIONS
*/

type CSSMAPPLETPCRLOPTIONS struct {
  Version uint `v8:"version"`
  CrlFlags uint `v8:"crlFlags"`
  CrlStore *CSSMDLDBHANDLEPTR `v8:"crlStore"`
}
