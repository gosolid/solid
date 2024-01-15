//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOHIDDeviceInterface122
*/

type IOHIDDeviceInterface122 struct {
  Reserved void `v8:"reserved"`
  QueryInterface func(...any) (any, error) `v8:"queryInterface"`
  AddRef func(...any) (any, error) `v8:"addRef"`
  Release func(...any) (any, error) `v8:"release"`
  CreateAsyncEventSource func(...any) (any, error) `v8:"createAsyncEventSource"`
  GetAsyncEventSource func(...any) (any, error) `v8:"getAsyncEventSource"`
  CreateAsyncPort func(...any) (any, error) `v8:"createAsyncPort"`
  GetAsyncPort func(...any) (any, error) `v8:"getAsyncPort"`
  Open func(...any) (any, error) `v8:"open"`
  Close func(...any) (any, error) `v8:"close"`
  SetRemovalCallback func(...any) (any, error) `v8:"setRemovalCallback"`
  GetElementValue func(...any) (any, error) `v8:"getElementValue"`
  SetElementValue func(...any) (any, error) `v8:"setElementValue"`
  QueryElementValue func(...any) (any, error) `v8:"queryElementValue"`
  StartAllQueues func(...any) (any, error) `v8:"startAllQueues"`
  StopAllQueues func(...any) (any, error) `v8:"stopAllQueues"`
  AllocQueue func(...any) (any, error) `v8:"allocQueue"`
  AllocOutputTransaction func(...any) (any, error) `v8:"allocOutputTransaction"`
  SetReport func(...any) (any, error) `v8:"setReport"`
  GetReport func(...any) (any, error) `v8:"getReport"`
  CopyMatchingElements func(...any) (any, error) `v8:"copyMatchingElements"`
  SetInterruptReportHandlerCallback func(...any) (any, error) `v8:"setInterruptReportHandlerCallback"`
}
