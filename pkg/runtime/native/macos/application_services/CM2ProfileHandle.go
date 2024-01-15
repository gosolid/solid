//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CM2ProfileHandle
*/

type CM2ProfileHandle struct {
  Header CM2Header `v8:"header"`
  TagTable CMTagElemTable `v8:"tagTable"`
  ElemData [1]byte `v8:"elemData"`
}
