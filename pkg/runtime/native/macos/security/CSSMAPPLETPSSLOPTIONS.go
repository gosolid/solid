//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_APPLE_TP_SSL_OPTIONS
*/

type CSSMAPPLETPSSLOPTIONS struct {
  Version uint `v8:"version"`
  ServerNameLen uint `v8:"serverNameLen"`
  ServerName byte `v8:"serverName"`
  Flags uint `v8:"flags"`
}
