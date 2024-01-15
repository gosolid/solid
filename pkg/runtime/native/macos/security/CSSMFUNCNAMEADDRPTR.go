//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_FUNC_NAME_ADDR_PTR
*/

type CSSMFUNCNAMEADDRPTR struct {
  Name [68]byte `v8:"name"`
  Address *func(...any) (any, error) `v8:"address"`
}
