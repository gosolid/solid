//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_APPLE_TP_SMIME_OPTIONS
*/

type CSSMAPPLETPSMIMEOPTIONS struct {
  Version uint `v8:"version"`
  IntendedUsage uint16 `v8:"intendedUsage"`
  SenderEmailLen uint `v8:"senderEmailLen"`
  SenderEmail byte `v8:"senderEmail"`
}
