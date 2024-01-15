//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_KR_WRAPPEDPRODUCT_INFO_PTR
*/

type CSSMKRWRAPPEDPRODUCTINFOPTR struct {
  StandardVersion CSSMVERSIONPTR `v8:"standardVersion"`
  StandardDescription [68]byte `v8:"standardDescription"`
  ProductVersion CSSMVERSIONPTR `v8:"productVersion"`
  ProductDescription [68]byte `v8:"productDescription"`
  ProductVendor [68]byte `v8:"productVendor"`
  ProductFlags uint `v8:"productFlags"`
}
