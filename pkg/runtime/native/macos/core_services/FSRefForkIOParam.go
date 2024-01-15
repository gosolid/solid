//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.FSRefForkIOParam
*/

type FSRefForkIOParam struct {
  QLink *QElem `v8:"qLink"`
  QType int16 `v8:"qType"`
  IoTrap int16 `v8:"ioTrap"`
  IoCmdAddr *byte `v8:"ioCmdAddr"`
  IoCompletion *func(...any) (any, error) `v8:"ioCompletion"`
  IoResult int16 `v8:"ioResult"`
  ParentRef FSRef `v8:"parentRef"`
  NameLength uint64 `v8:"nameLength"`
  Name uint16 `v8:"name"`
  WhichInfo uint `v8:"whichInfo"`
  CatInfo FSCatalogInfo `v8:"catInfo"`
  ForkNameLength uint64 `v8:"forkNameLength"`
  ForkName uint16 `v8:"forkName"`
  Permissions byte `v8:"permissions"`
  Reserved1 byte `v8:"reserved1"`
  ForkRefNum int `v8:"forkRefNum"`
  NewRef FSRef `v8:"newRef"`
}
