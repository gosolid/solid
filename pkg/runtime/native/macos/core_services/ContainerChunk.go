//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ContainerChunk
*/

type ContainerChunk struct {
  CkID uint `v8:"ckID"`
  CkSize int `v8:"ckSize"`
  FormType uint `v8:"formType"`
}
