//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSUTab
*/

type ATSUTab struct {
  TabPosition int `v8:"tabPosition"`
  TabType uint16 `v8:"tabType"`
}
