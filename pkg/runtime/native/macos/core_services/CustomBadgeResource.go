//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.CustomBadgeResource
*/

type CustomBadgeResource struct {
  Version int16 `v8:"version"`
  CustomBadgeResourceID int16 `v8:"customBadgeResourceID"`
  CustomBadgeType uint `v8:"customBadgeType"`
  CustomBadgeCreator uint `v8:"customBadgeCreator"`
  WindowBadgeType uint `v8:"windowBadgeType"`
  WindowBadgeCreator uint `v8:"windowBadgeCreator"`
  OverrideType uint `v8:"overrideType"`
  OverrideCreator uint `v8:"overrideCreator"`
}
