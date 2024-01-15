//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.LocalDateTime
*/

type LocalDateTime struct {
  HighSeconds uint16 `v8:"highSeconds"`
  LowSeconds uint `v8:"lowSeconds"`
  Fraction uint16 `v8:"fraction"`
}
