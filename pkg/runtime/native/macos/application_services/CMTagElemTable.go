//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMTagElemTable
*/

type CMTagElemTable struct {
  Count uint `v8:"count"`
  TagList [1]CMTagRecord `v8:"tagList"`
}
