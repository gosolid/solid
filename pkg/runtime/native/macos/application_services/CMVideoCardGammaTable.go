//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMVideoCardGammaTable
*/

type CMVideoCardGammaTable struct {
  Channels uint16 `v8:"channels"`
  EntryCount uint16 `v8:"entryCount"`
  EntrySize uint16 `v8:"entrySize"`
  Data [1]byte `v8:"data"`
}
