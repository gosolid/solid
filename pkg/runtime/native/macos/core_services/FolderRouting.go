//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.FolderRouting
*/

type FolderRouting struct {
  DescSize int64 `v8:"descSize"`
  FileType uint `v8:"fileType"`
  RouteFromFolder uint `v8:"routeFromFolder"`
  RouteToFolder uint `v8:"routeToFolder"`
  Flags uint `v8:"flags"`
}
