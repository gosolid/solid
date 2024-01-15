//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.QElem
*/

type QElem struct {
  QLink QElem `v8:"qLink"`
  QType int16 `v8:"qType"`
  QData [1]int16 `v8:"qData"`
}
