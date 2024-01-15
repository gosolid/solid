//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSFlatDataStyleRunDataHeader
*/

type ATSFlatDataStyleRunDataHeader struct {
  NumberOfStyleRuns uint `v8:"numberOfStyleRuns"`
  StyleRunArray [1]ATSUStyleRunInfo `v8:"styleRunArray"`
}
