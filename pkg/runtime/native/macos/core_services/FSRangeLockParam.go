//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.FSRangeLockParam
*/

type FSRangeLockParam struct {
  QLink *QElem `v8:"qLink"`
  QType int16 `v8:"qType"`
  IoTrap int16 `v8:"ioTrap"`
  IoCmdAddr *byte `v8:"ioCmdAddr"`
  IoCompletion *func(...any) (any, error) `v8:"ioCompletion"`
  IoResult int16 `v8:"ioResult"`
  ForkRefNum int `v8:"forkRefNum"`
  RequestCount uint64 `v8:"requestCount"`
  PositionMode uint16 `v8:"positionMode"`
  PositionOffset int64 `v8:"positionOffset"`
  RangeStart uint64 `v8:"rangeStart"`
}
