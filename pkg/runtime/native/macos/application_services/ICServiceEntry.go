//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ICServiceEntry
*/

type ICServiceEntry struct {
  Name [256]byte `v8:"name"`
  Port int16 `v8:"port"`
  Flags int16 `v8:"flags"`
}
