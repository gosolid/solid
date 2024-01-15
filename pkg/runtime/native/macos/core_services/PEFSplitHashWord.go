//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.PEFSplitHashWord
*/

type PEFSplitHashWord struct {
  NameLength uint16 `v8:"nameLength"`
  HashValue uint16 `v8:"hashValue"`
}
