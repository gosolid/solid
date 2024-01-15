//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.DeferredTask
*/

type DeferredTask struct {
  QLink *QElem `v8:"qLink"`
  QType int16 `v8:"qType"`
  DtFlags int16 `v8:"dtFlags"`
  DtAddr *func(...any) (any, error) `v8:"dtAddr"`
  DtParam int64 `v8:"dtParam"`
  DtReserved int64 `v8:"dtReserved"`
}
