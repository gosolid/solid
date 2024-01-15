//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOFireWireSBP2LibMgmtORBInterface
*/

type IOFireWireSBP2LibMgmtORBInterface struct {
  Reserved void `v8:"reserved"`
  QueryInterface func(...any) (any, error) `v8:"queryInterface"`
  AddRef func(...any) (any, error) `v8:"addRef"`
  Release func(...any) (any, error) `v8:"release"`
  Version uint16 `v8:"version"`
  Revision uint16 `v8:"revision"`
  SubmitORB func(...any) (any, error) `v8:"submitORB"`
  SetORBCompleteCallback func(...any) (any, error) `v8:"setORBCompleteCallback"`
  SetRefCon func(...any) (any, error) `v8:"setRefCon"`
  GetRefCon func(...any) (any, error) `v8:"getRefCon"`
  SetCommandFunction func(...any) (any, error) `v8:"setCommandFunction"`
  SetManageeORB func(...any) (any, error) `v8:"setManageeORB"`
  SetManageeLogin func(...any) (any, error) `v8:"setManageeLogin"`
  SetResponseBuffer func(...any) (any, error) `v8:"setResponseBuffer"`
}
