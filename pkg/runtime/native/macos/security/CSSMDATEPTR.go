//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_DATE_PTR
*/

type CSSMDATEPTR struct {
  Year [4]byte `v8:"year"`
  Month [2]byte `v8:"month"`
  Day [2]byte `v8:"day"`
}
