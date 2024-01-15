//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct CoreServices.FSForkIOParam
*/

type FSForkIOParam struct {
  QLink *QElem `v8:"qLink"`
  QType int16 `v8:"qType"`
  IoTrap int16 `v8:"ioTrap"`
  IoCmdAddr *byte `v8:"ioCmdAddr"`
  IoCompletion *func(...any) (any, error) `v8:"ioCompletion"`
  IoResult int16 `v8:"ioResult"`
  Reserved1 void `v8:"reserved1"`
  Reserved2 int16 `v8:"reserved2"`
  ForkRefNum int `v8:"forkRefNum"`
  Reserved3 byte `v8:"reserved3"`
  Permissions byte `v8:"permissions"`
  Ref FSRef `v8:"ref"`
  Buffer *byte `v8:"buffer"`
  RequestCount uint `v8:"requestCount"`
  ActualCount uint `v8:"actualCount"`
  PositionMode uint16 `v8:"positionMode"`
  PositionOffset int64 `v8:"positionOffset"`
  AllocationFlags uint16 `v8:"allocationFlags"`
  AllocationAmount uint64 `v8:"allocationAmount"`
  ForkNameLength uint64 `v8:"forkNameLength"`
  ForkName uint16 `v8:"forkName"`
  ForkIterator CatPositionRec `v8:"forkIterator"`
  OutForkName objc.HFSUniStr255 `v8:"outForkName"`
}
