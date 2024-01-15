//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOHIDDeviceDeviceInterface
*/

type IOHIDDeviceDeviceInterface struct {
  Reserved void `v8:"reserved"`
  QueryInterface func(...any) (any, error) `v8:"queryInterface"`
  AddRef func(...any) (any, error) `v8:"addRef"`
  Release func(...any) (any, error) `v8:"release"`
  Open func(...any) (any, error) `v8:"open"`
  Close func(...any) (any, error) `v8:"close"`
  GetProperty func(...any) (any, error) `v8:"getProperty"`
  SetProperty func(...any) (any, error) `v8:"setProperty"`
  GetAsyncEventSource func(...any) (any, error) `v8:"getAsyncEventSource"`
  CopyMatchingElements func(...any) (any, error) `v8:"copyMatchingElements"`
  SetValue func(...any) (any, error) `v8:"setValue"`
  GetValue func(...any) (any, error) `v8:"getValue"`
  SetInputReportCallback func(...any) (any, error) `v8:"setInputReportCallback"`
  SetReport func(...any) (any, error) `v8:"setReport"`
  GetReport func(...any) (any, error) `v8:"getReport"`
}
