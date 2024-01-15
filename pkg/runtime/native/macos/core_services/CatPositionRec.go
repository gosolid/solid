//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.CatPositionRec
*/

type CatPositionRec struct {
  Initialize int `v8:"initialize"`
  Priv [6]int16 `v8:"priv"`
}
