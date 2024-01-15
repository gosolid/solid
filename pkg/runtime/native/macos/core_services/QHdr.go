//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.QHdr
*/

type QHdr struct {
  QFlags int16 `v8:"qFlags"`
  QHead *QElem `v8:"qHead"`
  QTail *QElem `v8:"qTail"`
}
