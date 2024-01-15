//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_SAMPLEGROUP_PTR
*/

type CSSMSAMPLEGROUPPTR struct {
  NumberOfSamples uint `v8:"numberOfSamples"`
  Samples CSSMSAMPLEPTR `v8:"samples"`
}
