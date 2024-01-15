//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_SPI_AC_FUNCS_PTR
*/

type CSSMSPIACFUNCSPTR struct {
  AuthCompute func(...any) (any, error) `v8:"authCompute"`
  PassThrough func(...any) (any, error) `v8:"passThrough"`
}
