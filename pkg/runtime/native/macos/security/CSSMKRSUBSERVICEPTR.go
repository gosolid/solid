//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_KRSUBSERVICE_PTR
*/

type CSSMKRSUBSERVICEPTR struct {
  SubServiceId uint `v8:"subServiceId"`
  Description byte `v8:"description"`
  WrappedProduct CSSMKRWRAPPEDPRODUCTINFOPTR `v8:"wrappedProduct"`
}
