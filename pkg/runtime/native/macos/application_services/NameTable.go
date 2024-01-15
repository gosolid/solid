//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.NameTable
*/

type NameTable struct {
  StringCount int16 `v8:"stringCount"`
  BaseFontName [256]byte `v8:"baseFontName"`
}
