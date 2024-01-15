//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ICMapEntry
*/

type ICMapEntry struct {
  TotalLength int16 `v8:"totalLength"`
  FixedLength int16 `v8:"fixedLength"`
  Version int16 `v8:"version"`
  FileType uint `v8:"fileType"`
  FileCreator uint `v8:"fileCreator"`
  PostCreator uint `v8:"postCreator"`
  Flags int `v8:"flags"`
  Extension [256]byte `v8:"extension"`
  CreatorAppName [256]byte `v8:"creatorAppName"`
  PostAppName [256]byte `v8:"postAppName"`
  MIMEType [256]byte `v8:"mIMEType"`
  EntryName [256]byte `v8:"entryName"`
}
