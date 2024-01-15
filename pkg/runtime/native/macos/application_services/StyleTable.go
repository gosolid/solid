//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.StyleTable
*/

type StyleTable struct {
  FontClass int16 `v8:"fontClass"`
  Offset int `v8:"offset"`
  Reserved int `v8:"reserved"`
  Indexes [48]byte `v8:"indexes"`
}
