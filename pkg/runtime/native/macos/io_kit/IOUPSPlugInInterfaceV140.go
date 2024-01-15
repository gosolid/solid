//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOUPSPlugInInterface_v140
*/

type IOUPSPlugInInterfaceV140 struct {
  Reserved void `v8:"reserved"`
  QueryInterface func(...any) (any, error) `v8:"queryInterface"`
  AddRef func(...any) (any, error) `v8:"addRef"`
  Release func(...any) (any, error) `v8:"release"`
  GetProperties func(...any) (any, error) `v8:"getProperties"`
  GetCapabilities func(...any) (any, error) `v8:"getCapabilities"`
  GetEvent func(...any) (any, error) `v8:"getEvent"`
  SetEventCallback func(...any) (any, error) `v8:"setEventCallback"`
  SendCommand func(...any) (any, error) `v8:"sendCommand"`
  CreateAsyncEventSource func(...any) (any, error) `v8:"createAsyncEventSource"`
}
