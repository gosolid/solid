//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.FSPermissionInfo
*/

type FSPermissionInfo struct {
  UserID uint `v8:"userID"`
  GroupID uint `v8:"groupID"`
  Reserved1 byte `v8:"reserved1"`
  UserAccess byte `v8:"userAccess"`
  Mode uint16 `v8:"mode"`
  FileSec *FSFileSecurity `v8:"fileSec"`
}
