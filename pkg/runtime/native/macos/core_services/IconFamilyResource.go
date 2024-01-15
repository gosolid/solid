//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.IconFamilyResource
*/

type IconFamilyResource struct {
  ResourceType uint `v8:"resourceType"`
  ResourceSize int `v8:"resourceSize"`
  Elements [1]IconFamilyElement `v8:"elements"`
}
