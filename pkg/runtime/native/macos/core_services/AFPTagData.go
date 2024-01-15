//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.AFPTagData
*/

type AFPTagData struct {
  FLength byte `v8:"fLength"`
  FType byte `v8:"fType"`
  FData [1]byte `v8:"fData"`
}
