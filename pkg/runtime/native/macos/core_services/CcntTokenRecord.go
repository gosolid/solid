//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ccntTokenRecord
*/

type CcntTokenRecord struct {
  TokenClass uint `v8:"tokenClass"`
  Token AEDesc `v8:"token"`
}
