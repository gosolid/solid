//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.AVLTreeStruct
*/

type AVLTreeStruct struct {
  Signature uint `v8:"signature"`
  PrivateStuff [8]uint64 `v8:"privateStuff"`
}
