//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.PhonemeInfo
*/

type PhonemeInfo struct {
  Opcode int16 `v8:"opcode"`
  PhStr [16]byte `v8:"phStr"`
  ExampleStr [32]byte `v8:"exampleStr"`
  HiliteStart int16 `v8:"hiliteStart"`
  HiliteEnd int16 `v8:"hiliteEnd"`
}
