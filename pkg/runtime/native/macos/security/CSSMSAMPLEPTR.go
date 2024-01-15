//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_SAMPLE_PTR
*/

type CSSMSAMPLEPTR struct {
  TypedSample CSSMLISTPTR `v8:"typedSample"`
  Verifier CSSMSUBSERVICEUIDPTR `v8:"verifier"`
}
