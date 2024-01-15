//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct CoreServices.FSRefParam
*/

type FSRefParam struct {
  QLink *QElem `v8:"qLink"`
  QType int16 `v8:"qType"`
  IoTrap int16 `v8:"ioTrap"`
  IoCmdAddr *byte `v8:"ioCmdAddr"`
  IoCompletion *func(...any) (any, error) `v8:"ioCompletion"`
  IoResult int16 `v8:"ioResult"`
  IoNamePtr *byte `v8:"ioNamePtr"`
  IoVRefNum int16 `v8:"ioVRefNum"`
  Reserved1 int16 `v8:"reserved1"`
  Reserved2 byte `v8:"reserved2"`
  Reserved3 byte `v8:"reserved3"`
  Ref FSRef `v8:"ref"`
  WhichInfo uint `v8:"whichInfo"`
  CatInfo FSCatalogInfo `v8:"catInfo"`
  NameLength uint64 `v8:"nameLength"`
  Name uint16 `v8:"name"`
  IoDirID uint `v8:"ioDirID"`
  Spec *FSSpec `v8:"spec"`
  ParentRef FSRef `v8:"parentRef"`
  NewRef FSRef `v8:"newRef"`
  TextEncodingHint uint `v8:"textEncodingHint"`
  OutName objc.HFSUniStr255 `v8:"outName"`
}
