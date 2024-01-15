//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_CERT_BUNDLE_HEADER_PTR
*/

type CSSMCERTBUNDLEHEADERPTR struct {
  BundleType uint `v8:"bundleType"`
  BundleEncoding uint `v8:"bundleEncoding"`
}
