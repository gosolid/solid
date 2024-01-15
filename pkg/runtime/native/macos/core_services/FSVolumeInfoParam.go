//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct CoreServices.FSVolumeInfoParam
*/

type FSVolumeInfoParam struct {
  QLink *QElem `v8:"qLink"`
  QType int16 `v8:"qType"`
  IoTrap int16 `v8:"ioTrap"`
  IoCmdAddr *byte `v8:"ioCmdAddr"`
  IoCompletion *func(...any) (any, error) `v8:"ioCompletion"`
  IoResult int16 `v8:"ioResult"`
  IoNamePtr *byte `v8:"ioNamePtr"`
  IoVRefNum int16 `v8:"ioVRefNum"`
  VolumeIndex uint `v8:"volumeIndex"`
  WhichInfo uint `v8:"whichInfo"`
  VolumeInfo FSVolumeInfo `v8:"volumeInfo"`
  VolumeName objc.HFSUniStr255 `v8:"volumeName"`
  Ref FSRef `v8:"ref"`
}
