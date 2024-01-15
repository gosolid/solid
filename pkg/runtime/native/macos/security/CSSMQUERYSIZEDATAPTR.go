//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_QUERY_SIZE_DATA_PTR
*/

type CSSMQUERYSIZEDATAPTR struct {
  SizeInputBlock uint `v8:"sizeInputBlock"`
  SizeOutputBlock uint `v8:"sizeOutputBlock"`
}
