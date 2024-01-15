//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TUPLEGROUP_PTR
*/

type CSSMTUPLEGROUPPTR struct {
  NumberOfTuples uint `v8:"numberOfTuples"`
  Tuples *CSSMTUPLEPTR `v8:"tuples"`
}
