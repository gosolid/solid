//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ExtendedFileInfo
*/

type ExtendedFileInfo struct {
  Reserved1 [4]int16 `v8:"reserved1"`
  ExtendedFinderFlags uint16 `v8:"extendedFinderFlags"`
  Reserved2 int16 `v8:"reserved2"`
  PutAwayFolderID int `v8:"putAwayFolderID"`
}
