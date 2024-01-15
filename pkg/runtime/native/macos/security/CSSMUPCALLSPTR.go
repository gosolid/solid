//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_UPCALLS_PTR
*/

type CSSMUPCALLSPTR struct {
  MallocFunc *func(...any) (any, error) `v8:"mallocFunc"`
  FreeFunc *func(...any) (any, error) `v8:"freeFunc"`
  ReallocFunc *func(...any) (any, error) `v8:"reallocFunc"`
  CallocFunc *func(...any) (any, error) `v8:"callocFunc"`
  CcToHandleFunc func(...any) (any, error) `v8:"ccToHandleFunc"`
  GetModuleInfoFunc func(...any) (any, error) `v8:"getModuleInfoFunc"`
}
