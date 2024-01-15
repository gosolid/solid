//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.OffPair
*/

type OffPair struct {
  OffFirst int16 `v8:"offFirst"`
  OffSecond int16 `v8:"offSecond"`
}
