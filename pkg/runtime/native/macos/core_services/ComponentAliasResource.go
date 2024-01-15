//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ComponentAliasResource
*/

type ComponentAliasResource struct {
  Cr ComponentResource `v8:"cr"`
  AliasCD ComponentDescription `v8:"aliasCD"`
}
