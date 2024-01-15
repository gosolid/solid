//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.GrafPort
*/

type GrafPort struct {
  Whatever [87]int16 `v8:"whatever"`
}
