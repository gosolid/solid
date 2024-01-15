//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_APPLE_TP_ACTION_DATA
*/

type CSSMAPPLETPACTIONDATA struct {
  Version uint `v8:"version"`
  ActionFlags uint `v8:"actionFlags"`
}
