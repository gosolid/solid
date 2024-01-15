//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_GUID_PTR
*/

type CSSMGUIDPTR struct {
  Data1 uint `v8:"data1"`
  Data2 uint16 `v8:"data2"`
  Data3 uint16 `v8:"data3"`
  Data4 [8]byte `v8:"data4"`
}
