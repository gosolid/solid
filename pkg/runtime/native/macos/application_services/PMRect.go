//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.PMRect
*/

type PMRect struct {
  Top float64 `v8:"top"`
  Left float64 `v8:"left"`
  Bottom float64 `v8:"bottom"`
  Right float64 `v8:"right"`
}
