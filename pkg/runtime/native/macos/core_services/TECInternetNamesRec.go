//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TECInternetNamesRec
*/

type TECInternetNamesRec struct {
  Count uint `v8:"count"`
  InternetNames TECInternetNameRec `v8:"internetNames"`
}
