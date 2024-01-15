//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOHIDDeviceQueueInterface
*/

type IOHIDDeviceQueueInterface struct {
  Reserved void `v8:"reserved"`
  QueryInterface func(...any) (any, error) `v8:"queryInterface"`
  AddRef func(...any) (any, error) `v8:"addRef"`
  Release func(...any) (any, error) `v8:"release"`
  GetAsyncEventSource func(...any) (any, error) `v8:"getAsyncEventSource"`
  SetDepth func(...any) (any, error) `v8:"setDepth"`
  GetDepth func(...any) (any, error) `v8:"getDepth"`
  AddElement func(...any) (any, error) `v8:"addElement"`
  RemoveElement func(...any) (any, error) `v8:"removeElement"`
  ContainsElement func(...any) (any, error) `v8:"containsElement"`
  Start func(...any) (any, error) `v8:"start"`
  Stop func(...any) (any, error) `v8:"stop"`
  SetValueAvailableCallback func(...any) (any, error) `v8:"setValueAvailableCallback"`
  CopyNextValue func(...any) (any, error) `v8:"copyNextValue"`
}
