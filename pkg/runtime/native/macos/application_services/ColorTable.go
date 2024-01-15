//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ColorTable
*/

type ColorTable struct {
  CtSeed int `v8:"ctSeed"`
  CtFlags int16 `v8:"ctFlags"`
  CtSize int16 `v8:"ctSize"`
  CtTable [1]ColorSpec `v8:"ctTable"`
}
