//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct CoreServices.FSForkCBInfoParam
*/

type FSForkCBInfoParam struct {
  QLink *QElem `v8:"qLink"`
  QType int16 `v8:"qType"`
  IoTrap int16 `v8:"ioTrap"`
  IoCmdAddr *byte `v8:"ioCmdAddr"`
  IoCompletion *func(...any) (any, error) `v8:"ioCompletion"`
  IoResult int16 `v8:"ioResult"`
  DesiredRefNum int `v8:"desiredRefNum"`
  VolumeRefNum int16 `v8:"volumeRefNum"`
  Iterator int `v8:"iterator"`
  ActualRefNum int16 `v8:"actualRefNum"`
  Ref FSRef `v8:"ref"`
  ForkInfo FSForkInfo `v8:"forkInfo"`
  ForkName objc.HFSUniStr255 `v8:"forkName"`
}
