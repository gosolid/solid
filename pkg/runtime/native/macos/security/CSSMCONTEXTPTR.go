//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_CONTEXT_PTR
*/

type CSSMCONTEXTPTR struct {
  ContextType uint `v8:"contextType"`
  AlgorithmType uint `v8:"algorithmType"`
  NumberOfAttributes uint `v8:"numberOfAttributes"`
  ContextAttributes *CSSMCONTEXTATTRIBUTEPTR `v8:"contextAttributes"`
  CSPHandle int64 `v8:"cSPHandle"`
  Privileged int `v8:"privileged"`
  EncryptionProhibited uint `v8:"encryptionProhibited"`
  WorkFactor uint `v8:"workFactor"`
  Reserved uint `v8:"reserved"`
}
