//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSUUnhighlightData
*/

type ATSUUnhighlightData struct {
  DataType uint `v8:"dataType"`
  UnhighlightData void `v8:"unhighlightData"`
}
