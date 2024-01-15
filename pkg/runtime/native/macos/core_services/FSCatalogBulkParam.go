//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct CoreServices.FSCatalogBulkParam
*/

type FSCatalogBulkParam struct {
  QLink *QElem `v8:"qLink"`
  QType int16 `v8:"qType"`
  IoTrap int16 `v8:"ioTrap"`
  IoCmdAddr *byte `v8:"ioCmdAddr"`
  IoCompletion *func(...any) (any, error) `v8:"ioCompletion"`
  IoResult int16 `v8:"ioResult"`
  ContainerChanged byte `v8:"containerChanged"`
  Reserved byte `v8:"reserved"`
  IteratorFlags uint `v8:"iteratorFlags"`
  Iterator *OpaqueFSIterator `v8:"iterator"`
  Container FSRef `v8:"container"`
  MaximumItems uint64 `v8:"maximumItems"`
  ActualItems uint64 `v8:"actualItems"`
  WhichInfo uint `v8:"whichInfo"`
  CatalogInfo FSCatalogInfo `v8:"catalogInfo"`
  Refs FSRef `v8:"refs"`
  Specs *FSSpec `v8:"specs"`
  Names objc.HFSUniStr255 `v8:"names"`
  SearchParams FSSearchParams `v8:"searchParams"`
}
