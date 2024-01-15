//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOHIDDeviceTransactionInterface
*/

type IOHIDDeviceTransactionInterface struct {
  Reserved void `v8:"reserved"`
  QueryInterface func(...any) (any, error) `v8:"queryInterface"`
  AddRef func(...any) (any, error) `v8:"addRef"`
  Release func(...any) (any, error) `v8:"release"`
  GetAsyncEventSource func(...any) (any, error) `v8:"getAsyncEventSource"`
  SetDirection func(...any) (any, error) `v8:"setDirection"`
  GetDirection func(...any) (any, error) `v8:"getDirection"`
  AddElement func(...any) (any, error) `v8:"addElement"`
  RemoveElement func(...any) (any, error) `v8:"removeElement"`
  ContainsElement func(...any) (any, error) `v8:"containsElement"`
  SetValue func(...any) (any, error) `v8:"setValue"`
  GetValue func(...any) (any, error) `v8:"getValue"`
  Commit func(...any) (any, error) `v8:"commit"`
  Clear func(...any) (any, error) `v8:"clear"`
}
