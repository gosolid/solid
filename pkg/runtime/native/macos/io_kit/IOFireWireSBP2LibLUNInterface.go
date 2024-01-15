//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOFireWireSBP2LibLUNInterface
*/

type IOFireWireSBP2LibLUNInterface struct {
  Reserved void `v8:"reserved"`
  QueryInterface func(...any) (any, error) `v8:"queryInterface"`
  AddRef func(...any) (any, error) `v8:"addRef"`
  Release func(...any) (any, error) `v8:"release"`
  Version uint16 `v8:"version"`
  Revision uint16 `v8:"revision"`
  Open func(...any) (any, error) `v8:"open"`
  OpenWithSessionRef func(...any) (any, error) `v8:"openWithSessionRef"`
  GetSessionRef func(...any) (any, error) `v8:"getSessionRef"`
  Close func(...any) (any, error) `v8:"close"`
  AddCallbackDispatcherToRunLoop func(...any) (any, error) `v8:"addCallbackDispatcherToRunLoop"`
  RemoveCallbackDispatcherFromRunLoop func(...any) (any, error) `v8:"removeCallbackDispatcherFromRunLoop"`
  SetMessageCallback func(...any) (any, error) `v8:"setMessageCallback"`
  SetRefCon func(...any) (any, error) `v8:"setRefCon"`
  GetRefCon func(...any) (any, error) `v8:"getRefCon"`
  CreateLogin func(...any) (any, error) `v8:"createLogin"`
  CreateMgmtORB func(...any) (any, error) `v8:"createMgmtORB"`
}
