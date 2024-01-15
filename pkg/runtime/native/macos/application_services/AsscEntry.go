//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.AsscEntry
*/

type AsscEntry struct {
  FontSize int16 `v8:"fontSize"`
  FontStyle int16 `v8:"fontStyle"`
  FontID int16 `v8:"fontID"`
}
