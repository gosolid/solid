//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_CERT_BUNDLE_PTR
*/

type CSSMCERTBUNDLEPTR struct {
  BundleHeader CSSMCERTBUNDLEHEADERPTR `v8:"bundleHeader"`
  Bundle SecAsn1Oid `v8:"bundle"`
}
