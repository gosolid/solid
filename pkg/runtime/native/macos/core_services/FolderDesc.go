//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.FolderDesc
*/

type FolderDesc struct {
  DescSize int64 `v8:"descSize"`
  FoldType uint `v8:"foldType"`
  Flags uint `v8:"flags"`
  FoldClass uint `v8:"foldClass"`
  FoldLocation uint `v8:"foldLocation"`
  BadgeSignature uint `v8:"badgeSignature"`
  BadgeType uint `v8:"badgeType"`
  Reserved uint `v8:"reserved"`
  Name [64]byte `v8:"name"`
}
