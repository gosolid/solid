//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_KR_NAME
*/

type CSSMKRNAME struct {
  Type byte `v8:"type"`
  Length byte `v8:"length"`
  Name byte `v8:"name"`
}
