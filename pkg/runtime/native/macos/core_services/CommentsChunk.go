//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.CommentsChunk
*/

type CommentsChunk struct {
  CkID uint `v8:"ckID"`
  CkSize int `v8:"ckSize"`
  NumComments uint16 `v8:"numComments"`
  Comments [1]Comment `v8:"comments"`
}
