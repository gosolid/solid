//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.AIFFLoop
*/

type AIFFLoop struct {
  PlayMode int16 `v8:"playMode"`
  BeginLoop int16 `v8:"beginLoop"`
  EndLoop int16 `v8:"endLoop"`
}
