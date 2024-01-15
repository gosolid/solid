//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_MEMORY_FUNCS_PTR
*/

type CSSMMEMORYFUNCSPTR struct {
  MallocFunc *func(...any) (any, error) `v8:"mallocFunc"`
  FreeFunc *func(...any) (any, error) `v8:"freeFunc"`
  ReallocFunc *func(...any) (any, error) `v8:"reallocFunc"`
  CallocFunc *func(...any) (any, error) `v8:"callocFunc"`
  AllocRef void `v8:"allocRef"`
}
